package main

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"testing"
)

var gbhttpaddr = "192.168.1.183:7777"

/*
*

	{
		"streamSource":"rtsp://:554/proxyGbLiveStream-MP2P/66012000001320021542/udp",
		"playType":"live",
		"transport":"udp",
		"proxyType":"RTSP",
		"remoteRecvIP":"192.168.2.11",
		"remoteRecvPort":10100
	}
*/
func TestMediaPlay(t *testing.T) {
	client := NewHTTPClient()
	req := map[string]interface{}{
		"audioPort":    40101,
		"cmd":          "start",
		"deviceId":     "33012000001320000011",
		"fromPlatform": "33012000002000000112",
		"ip":           "192.168.2.11",
		"linkMode":     "udp",
		"proto":        "GB28181-2016",
		"srId":         "8a2888db7771d63494b9",
		"videoPort":    10100,
	}
	response, err := client.Put("http://"+gbhttpaddr+"/v1/ipc/video/play", req)
	if err != nil {
		panic("err:" + err.Error())
	}
	mp := make(map[string]interface{})
	if err := jsoniter.Unmarshal(response, &mp); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("INVITE data:", mp["data"])

	// ack
	req["cmd"] = "ack"
	response2, err := client.Put("http://"+gbhttpaddr+"/v1/ipc/video/play", req)
	if err != nil {
		panic("err:" + err.Error())
	}
	mp = make(map[string]interface{})
	if err := jsoniter.Unmarshal(response2, &mp); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ACK data:", mp["data"])

}

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}

// Get 发送GET请求
func (c *HTTPClient) Get(url string) ([]byte, error) {
	//params := url.Values{}
	//params.Add("key1", "value1")
	//params.Add("key2", "value2")
	//
	//// 将查询参数添加到URL
	//rawQuery := params.Encode()
	//fullURL := baseURL + "?" + rawQuery
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Post 发送POST请求，支持JSON格式的数据
func (c *HTTPClient) Post(url string, data interface{}) ([]byte, error) {
	jsonData, err := jsoniter.Marshal(data)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonData))
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("http.NewRequest:", err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println("c.client.Do:", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Put 发送PUT请求，支持JSON格式的数据
func (c *HTTPClient) Put(url string, data interface{}) ([]byte, error) {
	jsonData, err := jsoniter.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SetupSessionReq struct {
	StreamSource string `json:"streamSource"` // 流源，rtsp流时给出rtsp url，国标流时给出ipc uuid
	PlayType     string `json:"playType"`     // 播放类型
	Transport    string `json:"transport"`    // 流传输方式active、passive、udp
}

type SetupLiveSessionReq struct {
	SetupSessionReq
	ProxyType string `json:"proxyType"` // 流协议类型，下级发来的流协议类型。
	DestIP    string `json:"destIP"`    // 收流端ip。传输模式为udp和active时提供
	DestPort  int    `json:"destPort"`  // 收流端口。传输模式为udp和active时提供
}

type SetupPlaybackSessionReq struct {
	SetupSessionReq
	RangeStart string `json:"rangeStart"` // 播放时间范围起始点
	RangeEnd   string `json:"rangeEnd"`   // 播放时间范围终点
	ProxyType  string `json:"proxyType"`  // 流协议类型，下级发来的流协议类型。
	UserAgent  string `json:"userAgent"`  // 用户代理
}

type PlayReq struct {
	Session string  `json:"session"`
	Range   string  `json:"range"`
	Scale   float64 `json:"scale"`
}

type PauseReq struct {
	Session string `json:"session"`
}

type KeepAliveReq struct {
	Session string `json:"session"`
}
