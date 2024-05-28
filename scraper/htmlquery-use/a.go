package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
)

func main() {
	doc, err := htmlquery.LoadURL(`C:\Users\dell\3D Objects\a.html`)
	if err != nil {
		panic(err.Error())
	}
	nodes, err := htmlquery.QueryAll(doc, "//a")
	if err != nil {
		panic(`not a valid XPath expression.`)
	}
	for _, node := range nodes {
		url := htmlquery.FindOne(node, "./a/@href")
		fmt.Println(url)
	}

}
