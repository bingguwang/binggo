package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

var osStop chan os.Signal

func init() {
	osStop := make(chan os.Signal, 1)
	signal.Notify(osStop, syscall.SIGTERM, syscall.SIGINT)
}
func main() {
	f := func() *gin.Engine {
		r := gin.Default()
		r.PUT("/gbChannel/subscribe", SubscribeGBChannelHandle)
		return r
	}

	httpSrv := &http.Server{
		Addr:    ":7788",
		Handler: f(),
	}

	var countSend int
	go func() {

		for {
			select {
			case <-osStop:
				fmt.Println("exit")
				return
			default:
				//v := rand.Intn(10) + 1
				v := 1 // 塞的比取的快
				time.Sleep(time.Duration(v) * time.Second)

				subscribeGBChannelManager.Range(func(key, value any) bool {
					countSend++
					fmt.Println("共发送消息:", countSend, "条")
					info := value.(*SubscribeChannelSubInfo)
					fmt.Println("to ", info.Ip)
					info.MessageChannel <- &GBChannelMessage{
						Tp:           "ipc",
						Uuid:         "123456",
						FromPlatform: "000000000000",
						Operate:      "update",
					}
					return true
				})
			}
		}

	}()

	if err := httpSrv.ListenAndServe(); err != nil {
		panic(err.Error())
		return
	}

}

var subscribeGBChannelManager = &sync.Map{}

func SubscribeGBChannelHandle(c *gin.Context) {
	clientIp := c.ClientIP()
	fmt.Println("clientIp:", clientIp)
	req := SubscribeGBChannelReq{} // 是否订阅
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if req.Port == 0 {
		c.JSON(http.StatusBadRequest, "port must be send")
		return
	}
	if len(req.Url) == 0 {
		c.JSON(http.StatusBadRequest, "url must be send")
		return
	}
	if len(req.KeepAliveUrl) == 0 {
		c.JSON(http.StatusBadRequest, "KeepAliveUrl must be send")
		return
	}
	if req.Expire == 0 { // 取消订阅
		value, ok := subscribeGBChannelManager.Load(clientIp)
		if !ok {
			c.JSON(http.StatusOK, "OK")
			return
		}
		info := value.(*SubscribeChannelSubInfo)
		fmt.Println("通知订阅")
		close(info.StopChan)
	} else { // 订阅
		if v, ok := subscribeGBChannelManager.Load(clientIp); ok {
			// 续租
			subInfo := v.(*SubscribeChannelSubInfo)
			subInfo.Expire = req.Expire
			subInfo.ResetChannel <- true
			c.JSON(http.StatusOK, "OK")
		} else {
			subInfo := NewSubscribeChannelSubInfo(clientIp, req.Port, req.Expire, req.Url, req.KeepAliveUrl)
			go subInfo.noticeLoopGBMessageChannel(osStop)
			subscribeGBChannelManager.Store(clientIp, subInfo)
			c.JSON(http.StatusOK, "OK")
			return
		}

	}
}

type SubscribeGBChannelReq struct {
	Url          string `json:"url"`          // 客户端提供的推送信息的接口
	KeepAliveUrl string `json:"keepAliveUrl"` // 客户端提供的维护心跳的接口
	Port         int64  `json:"port"`
	Expire       int64  `json:"expire"` // 订阅时长，单位s
}

func NewSubscribeChannelSubInfo(ip string, port, expire int64, url, keepaliveUrl string) *SubscribeChannelSubInfo {
	resp := &SubscribeChannelSubInfo{
		Ip:             ip,
		Port:           port,
		StopChan:       make(chan struct{}),
		KeepAliveUrl:   keepaliveUrl,
		NotifyUrl:      url,
		Expire:         expire,
		MessageChannel: make(chan *GBChannelMessage, 100),
		ResetChannel:   make(chan bool),
	}
	return resp
}

type SubscribeChannelSubInfo struct {
	Ip             string
	Port           int64
	NotifyUrl      string
	KeepAliveUrl   string
	Expire         int64
	StopChan       chan struct{} // 订阅存活信号
	MessageChannel chan *GBChannelMessage
	Lock           sync.Mutex
	Timer          *time.Timer // 租约定时器
	ResetChannel   chan bool   // 续租通道
}

type GBChannelMessage struct {
	Tp           string // 组织还是ipc
	Uuid         string
	FromPlatform string
	Operate      string
}

func (subInfo *SubscribeChannelSubInfo) resetTimer(newExpire int64) {
	subInfo.Lock.Lock()
	defer subInfo.Lock.Unlock()
	if !subInfo.Timer.Stop() {
		<-subInfo.Timer.C // 清理已经触发的定时器
	}
	subInfo.Timer.Reset(time.Duration(newExpire) * time.Second)
}

func (subInfo *SubscribeChannelSubInfo) noticeLoopGBMessageChannel(osStopChan chan os.Signal) {

	subInfo.Lock.Lock()
	subInfo.Timer = time.NewTimer(time.Duration(subInfo.Expire) * time.Second)
	subInfo.Lock.Unlock()
	defer subInfo.Timer.Stop()

	// 租期维护
	go func() {
		for {
			select {
			case <-subInfo.StopChan:
				subscribeGBChannelManager.Delete(subInfo.Ip)
				fmt.Println("订阅结束")
				return
			case <-subInfo.Timer.C:
				fmt.Println("租期到")
				subscribeGBChannelManager.Delete(subInfo.Ip)
				close(subInfo.StopChan)
				return
			case v := <-subInfo.ResetChannel:
				if v { // 续租
					fmt.Println("续租,新的expire是", subInfo.Expire)
					subInfo.resetTimer(subInfo.Expire)
				}
			}
		}
	}()

	// 消息发布
	for {
		//国标ipc和组织变化则http推送于订阅者
		select {
		case _, ok := <-osStopChan:
			if !ok {
				//通知关闭
				return
			}
		case <-subInfo.StopChan:
			fmt.Println("租期到, 结束向此客户推送变化:", subInfo.Ip)
			return
		case v := <-subInfo.MessageChannel: // 有变化了通知给客户
			subInfo.PushChangeMessageToGBChannel(v)
		}
	}
}

// 推送变更消息给客户端
func (subinfo *SubscribeChannelSubInfo) PushChangeMessageToGBChannel(data *GBChannelMessage) {
	url := fmt.Sprintf("http://%s:%v%s", subinfo.Ip, subinfo.Port, subinfo.NotifyUrl)
	method := "POST"

	marshal, _ := json.Marshal(data)
	payload := strings.NewReader(string(marshal))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", "Token "+gateway.TrinetLogin(ip, port))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

}
