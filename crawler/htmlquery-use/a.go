package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
)

/**
基础语法:
/			从该节点的子元素选取
//			从该节点的子孙元素选取
*			通配符
nodename	选取此节点的所有子节点
…			选取当前节点的父节点
@			选取属性

其实可以直接F12在页面的console里进行练习
输入相关的xpath表达式就可以获取对应的元素结果
$x("/")		整个页面
$x("/*")	整个页面的所有元素
$x("//*")	整个页面的所有元素
$x("//div") 整个页面的所有的 div 标签节点
$x('//*[@id="site-logo"]')		整个页面的所有元素里，获取 id 属性为 site-logo 的节点
$x('//*[@id="site-logo"]/..') 	整个页面的所有元素里，获取 id 属性为 site-logo 的节点的父节点
$x("//*[@id='ember21']//li") 	整个页面的所有元素里，获取 id 属性为 ember21此节点下的所有的li元素
$x("//*[@id='ember21']//li[1]") 整个页面的所有元素里，获取 id 属性为 ember21此节点下的第一个li元素
$x("(//div)[last()]") // 获取最后一个div标签节点,[last()表示选取最后一个
[@属性名='属性值' and @属性名='属性值']	与关系
[@属性名='属性值' or @属性名='属性值']	或关系
[text()='文本信息']	根据文本信息定位
[contains(text(),'文本信息')]	根据文本信息包含定位

*/

func main() {
	//doc, err := htmlquery.LoadDoc(`./aaa.html`)
	doc, err := htmlquery.LoadDoc(`E:\project\binggo\crawler\htmlquery-use\aaa.html`)
	if err != nil {
		panic(err.Error())
	}
	nodes, err := htmlquery.QueryAll(doc, "//a") // 整个页面的所有a标签节点
	if err != nil {
		panic(`not a valid XPath expression.`)
	}
	fmt.Println("nodes个数:", len(nodes))
	var i = 0
	for _, node := range nodes {
		text := htmlquery.InnerText(node)
		fmt.Println("a标签的文本值:", text)
		urlNode := htmlquery.FindOne(node, "./@href") // a标签节点的href元素
		i++
		if urlNode != nil {
			url := htmlquery.SelectAttr(urlNode, "href")
			fmt.Println("超链接", i, " : ", url)
		}
	}
}
