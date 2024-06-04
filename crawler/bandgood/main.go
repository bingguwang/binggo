package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"log"
	"regexp"
	"strings"
)

func main() {
	fmt.Println()
	testcollyxpath()
}

/*
Collector对象接受多种回调方法，有不同的作用，按调用顺序我列出来：
OnRequest。请求前
OnError。请求过程中发生错误
OnResponse。收到响应后
OnHTML。如果收到的响应内容是HTML调用它。
OnXML。如果收到的响应内容是XML 调用它。写爬虫基本用不到，所以上面我没有使用它。
OnScraped。在OnXML/OnHTML回调完成后调用。不过官网写的是Called after OnXML callbacks，实际上对于OnHTML也有效，大家可以注意一下。
*/
//https://www.banggood.com/discover-new.html?utmid=22238&rmmds=home-new&bid=35141&last_spm=1a981.Homepage.0001127287.0001779.80ee5f9e5d7b4ed59cec9ca8905b60ee
func testcollydom() {
	//创建新的采集器
	c := colly.NewCollector(
		//这次在colly.NewCollector里面加了一项colly.Async(true)，表示抓取是异步的
		colly.Async(true),
		//colly.IgnoreRobotsTxt(), // 忽略robots.txt
		//模拟浏览器
		colly.UserAgent("Mozilla / 5.0(Windows NT 10.0; Win64; x64) AppleWebKit / 537.36(KHTML, like Gecko) Chrome / 80.0.3987.122  Safari / 537.36"),
		//colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)
	//限制采集规则
	//在Colly里面非常方便控制并发度，只抓取符合某个(些)规则的URLS，有一句c.Limit(&colly.LimitRule{DomainGlob: "*.douban.*", Parallelism: 5})，表示限制只抓取域名是douban(域名后缀和二级域名不限制)的地址，当然还支持正则匹配某些符合的 URLS，具体的可以看官方文档。
	c.Limit(&colly.LimitRule{DomainGlob: "*.banggood.*", Parallelism: 5})
	/*
		另外Limit方法中也限制了并发是5。为什么要控制并发度呢？因为抓取的瓶颈往往来自对方网站的抓取频率的限制，如果在一段时间内达到某个抓取频率很容易被封，所以我们要控制抓取的频率。另外为了不给对方网站带来额外的压力和资源消耗，也应该控制你的抓取机制。
	*/

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML(".hd", func(e *colly.HTMLElement) {
		log.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
			strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.banggood.com/discover-new.html?utmid=22238&rmmds=home-new&bid=35141&last_spm=1a981.Homepage.0001127287.0001779.80ee5f9e5d7b4ed59cec9ca8905b60ee")
	c.Wait()

}

func testcollyxpath() {
	//创建新的采集器
	c := colly.NewCollector(
		//这次在colly.NewCollector里面加了一项colly.Async(true)，表示抓取是异步的
		colly.Async(true),
		//模拟浏览器
		colly.UserAgent("Mozilla / 5.0(Windows NT 10.0; Win64; x64) AppleWebKit / 537.36(KHTML, like Gecko) Chrome / 80.0.3987.122  Safari / 537.36"),
		//colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
		//最大深度2
		// 如果你将 MaxDepth 设置为 3，那么爬虫将只跟踪从初始页面开始的链接，直到深度为3的层级。
		//也就是说，爬虫将只访问初始页面上的链接、初始页面上链接的页面、以及这些页面上的链接的页面。超过深度3的链接将被忽略。
		//有的页面是有循环链接的，不限制深度可能会一直爬下去
		colly.MaxDepth(2),
	)
	//限制采集规格
	c.Limit(&colly.LimitRule{DomainGlob: "*.banggood.*", Parallelism: 5})
	/**
	type LimitRule struct {
		// 域名正则表达式。：用于匹配域名的正则表达式。只有匹配到的域名才会受到该规则的限制。
		例如，如果你想限制爬虫访问以 "example.com" 开头的所有子域名，
		你可以设置 DomainRegexp 为 ^https?://(.*\.)?example\.com/.*$。
		DomainRegexp string
		// 域名通配符。用于匹配域名的通配符模式。与正则表达式类似，只有匹配到的域名才会受到该规则的限制。例如，如果你想限制爬虫访问以 "example.com" 开头的所有子域名，你可以设置 DomainGlob 为 *.example.com/*
		DomainGlob string
		// 例如，设置 Delay 为 5 * time.Second 将使爬虫在发送两个请求之间等待 5 秒钟。
		Delay time.Duration
		// 额外的随机化延迟，被添加到 Delay 中。这有助于模拟真实用户的行为，以防止被目标网站识别为爬虫。
		例如，如果你设置 Delay 为 10 * time.Second，而 RandomDelay 为 5 * time.Second，
		则爬虫在发送请求之前将等待 10 到 15 秒之间的随机时间。
		RandomDelay time.Duration
		// 例如，如果设置 Parallelism 为 3，则最多会同时发送 3 个请求到匹配该规则的域名。
		Parallelism    int
	}
	*/
	//请求前
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	//出现错误
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// 在收到 完整的HTTP 响应时执行的回调函数。
	c.OnResponse(func(r *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body))) //将 HTTP 响应的内容解析为 HTML 文档
		if err != nil {
			log.Fatal(err)
		}
		//使用 XPath 查询对其进行处理。它使用了 htmlquery 包来解析 HTML。
		// 它查找了所有 class 为 "hd" 的 <div> 元素，这些元素通常包含了电影的标题和链接等信息。
		nodes := htmlquery.Find(doc, `//ol[@class="grid_view"]/li//div[@class="hd"]`)
		for _, node := range nodes {
			//在每个节点中查找电影链接和标题
			url := htmlquery.FindOne(node, "./a/@href")
			title := htmlquery.FindOne(node, `.//span[@class="title"]/text()`)
			log.Println(strings.Split(htmlquery.InnerText(url), "/")[4],
				htmlquery.InnerText(title))
		}
	})
	//OnHTML指定了当 colly 在 HTML 页面中找到符合指定 CSS 选择器的链接时要执行的回调函数。
	//a[href]指定了所有带有 href 属性的 <a> 标签,每个符合的元素只会调用一次回调函数
	//因为最大深度设置2，
	//当前第一级 html里的 每个a标签都会回调访问
	c.OnHTML("a[href]", func(e *colly.HTMLElement) { // e是一个HTML元素
		link := e.Attr("href") // 获取此HTML元素的href属性
		fmt.Println("link:", link)

		// 查找行首以 ?start=0&filter= 的字符串（非贪婪模式）
		reg := regexp.MustCompile(`(?U)^\?start=(\d+)&filter=`)
		regMatch := reg.FindAllString(link, -1)
		//如果找的到的话
		if len(regMatch) > 0 {

			link = "https://movie.douban.com/top250" + regMatch[0]
			//访问该链接
			e.Request.Visit(link)
		}

		// Visit link found on page
	})

	//结束
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	//采集开始
	c.Visit("https://www.banggood.com/discover-new.html?utmid=22238&rmmds=home-new&bid=35141&last_spm=1a981.Homepage.0001127287.0001779.80ee5f9e5d7b4ed59cec9ca8905b60ee")
	c.Wait()

}
