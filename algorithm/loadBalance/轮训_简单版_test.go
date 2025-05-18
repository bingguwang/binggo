package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	loadBalancer := NewLoadBalancer([]*server{
		{uuid: "xxxxxx"},
		{uuid: "yyyyyy"},
	})

	for i := 0; i < 10; i++ {
		srv := loadBalancer.getloadBalancer()
		fmt.Println("使用:", srv.uuid)
	}
}

type server struct {
	uuid string
}
type loader struct {
	srvlst   []*server
	curindex int
}

// NewLoadBalancer 创建一个新的负载均衡器
func NewLoadBalancer(servers []*server) *loader {
	return &loader{
		srvlst:   servers,
		curindex: 0,
	}
}

func (l *loader) getloadBalancer() *server {
	var res *server
	if len(l.srvlst) == 0 {
		return nil
	}
	res = l.srvlst[l.curindex]
	l.curindex = (l.curindex + 1) % len(l.srvlst)
	return res
}
