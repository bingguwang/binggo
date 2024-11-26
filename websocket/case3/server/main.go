package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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
	if !req.Subscribe { // 取消订阅
		value, ok := subscribeGBChannelManager.Load(clientIp)
		if !ok {
			c.JSON(http.StatusOK, "OK")
			return
		}
		info := value.(*SubscribeChannelSubInfo)
		fmt.Println("通知订阅:", info.Chan)
		close(info.Chan)
	} else { // 订阅
		if _, ok := subscribeGBChannelManager.Load(clientIp); ok {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("客户端[%s]已订阅过", clientIp))
			return
		}
		subInfo := NewSubscribeChannelSubInfo(clientIp, strconv.Itoa(req.Port), req.Url, req.KeepAliveUrl)
		go subInfo.noticeLoopGBMessageChannel(osStop)
		subscribeGBChannelManager.Store(clientIp, subInfo)
		c.JSON(http.StatusOK, "OK")
		return
	}
}

type SubscribeGBChannelReq struct {
	Url          string `json:"url"` // 客户端提供的推送信息的接口
	KeepAliveUrl string `json:"keepAliveUrl"`
	Port         int    `json:"port"`
	Subscribe    bool   `json:"subscribe"`
}

func NewSubscribeChannelSubInfo(ip, port, url, keepaliveUrl string) *SubscribeChannelSubInfo {
	resp := &SubscribeChannelSubInfo{
		Ip:             ip,
		Port:           port,
		Url:            url,
		KeepAliveUrl:   keepaliveUrl,
		Chan:           make(chan struct{}),
		MessageChannel: make(chan *GBChannelMessage, 2),
	}
	return resp
}

type SubscribeChannelSubInfo struct {
	Ip             string
	Port           string
	Url            string
	KeepAliveUrl   string
	Chan           chan struct{} // 订阅存活信号
	MessageChannel chan *GBChannelMessage
}

type GBChannelMessage struct {
	Tp           string // 组织还是ipc
	Uuid         string
	FromPlatform string
	Operate      string
}

func (subInfo *SubscribeChannelSubInfo) noticeLoopGBMessageChannel(osStopChan chan os.Signal) {

	ticker := time.NewTicker(5 * time.Second)
	go func() {
		fmt.Println("开启心跳监听")
		for {
			select {
			case <-subInfo.Chan:
				fmt.Println("关闭订阅")
				return
			case <-osStopChan:
				fmt.Println("exit noticeLoopGBMessageChannel")
				return
			case <-ticker.C:
				if err := subInfo.KeepAlive(); err != nil {
					fmt.Println("心跳停止, 关闭订阅")
					close(subInfo.Chan)
					return
				}
			}
		}
	}()

	for {
		//国标ipc和组织变化则http推送于订阅者
		select {
		case _, ok := <-osStopChan:
			if !ok {
				//通知关闭
				return
			}
		case <-subInfo.Chan:
			fmt.Println("心跳停止结束向此客户推送变化:", subInfo.Ip)
			return
		case <-osStopChan:
			fmt.Println("exit push")
			return
		case v := <-subInfo.MessageChannel: // 有变化了通知给客户
			time.Sleep(3 * time.Second)
			subInfo.PushChangeMessageToGBChannel(v)
		}
	}
}

// 推送变更消息给客户端
func (subinfo *SubscribeChannelSubInfo) PushChangeMessageToGBChannel(data *GBChannelMessage) {
	url := "http://" + subinfo.Ip + ":" + subinfo.Port + subinfo.Url
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

func (subInfo *SubscribeChannelSubInfo) KeepAlive() error {
	fmt.Println(subInfo.Ip, " ", subInfo.Port)
	url := "http://" + subInfo.Ip + ":" + subInfo.Port + subInfo.KeepAliveUrl
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
