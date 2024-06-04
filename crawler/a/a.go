package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/http"
	"strings"
)

// #要爬取的网页链接
// var baseurl = "https://movie.douban.com/top250?start="
var baseurl = "https://www.baidu.com"

func main() {

	//datalist := getData(baseurl)

	//Collector 管理网络通信，并负责在 collector 作业运行时执行附加的回调
	c := colly.NewCollector()

	// 下面都是回调函数
	// Find and visit all links
	// 请求前执行的回调
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//在请求区间出错时调用
	c.OnError(func(response *colly.Response, err error) {

	})

	// 收到响应后的执行回调
	c.OnResponse(func(response *colly.Response) {

	})

	// 如果收到的内容是 HTML。则在OnResponse之后执行回调
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	// 如果收到的内容是 HTML 或 XML, 在OnHTML后执行此回调
	c.OnXML("", func(element *colly.XMLElement) {

	})

	//OnXML 回调后调用
	c.OnScraped(func(response *colly.Response) {

	})

	//c.Visit("http://go-colly.org/")
	c.Visit(baseurl)
}

// # 得到指定一个URL的网页内容
func askURL(url string) {
	client := &http.Client{}
	method := "GET"
	payload := strings.NewReader(``)
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 模拟浏览器头部信息，向豆瓣服务器发送消息
	// 用户代理，表示告诉豆瓣服务器，我们是什么类型的机器、浏览器（本质上是告诉浏览器，我们可以接收什么水平的文件内容）
	request.Header.Add("User-Agent", "Mozilla / 5.0(Windows NT 10.0; Win64; x64) AppleWebKit / 537.36(KHTML, like Gecko) Chrome / 80.0.3987.122  Safari / 537.36")

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

}
