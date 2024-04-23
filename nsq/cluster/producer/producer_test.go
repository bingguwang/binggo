package main

import (
	"fmt"
	"testing"
)

func TestGetTopicNodeAddrSet(t *testing.T) {
	arr := getTopicNodeAddrSet("queue1")
	fmt.Println(arr)
}
