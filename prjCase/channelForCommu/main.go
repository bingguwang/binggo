package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Superior struct {
	regChan              chan string // 重注册信号
	userAgent            *userAgent
	regStatus            string
	state                string
	heartBeatChan        chan string // 用于关闭 heartCheck 协程
	makeSuperiorDownChan chan string // 用于关闭 MakeSuperiorDown 协程

	relive int // 假设上级在下级重新注册n次后
}

func NewSuperior() *Superior {
	return &Superior{
		regChan:       make(chan string),
		heartBeatChan: make(chan string),
		state:         "online",
		userAgent:     &userAgent{},
	}
}

type userAgent struct {
}

func main() {
	superior := NewSuperior()
	// 注册
	go superior.register()

	// 检测上级心跳
	go superior.heartCheck()

	// 上级根据指定规则下线
	go superior.MakeSuperiorDown()

	defer func() {
		close(superior.heartBeatChan)
	}()
	osStopChan := make(chan os.Signal, 1)
	signal.Notify(osStopChan, syscall.SIGTERM, syscall.SIGINT)
	<-osStopChan
}

// 向上级注册
func (s *Superior) register() {
	for {
		if err := s.userAgent.register(); err == nil {
			fmt.Println("注册成功")
			s.regStatus = "SUCCESS"
			break
		} else {
			fmt.Println("注册失败，5s后重新注册")
		}

		select {
		case <-s.regChan:
			return
		case <-time.After(5 * time.Second):
		}
	}

	isReReg := false
	for {
		if isReReg {
			for {
				if err := s.userAgent.register(); err == nil {
					fmt.Println("注册成功")
					s.regStatus = "SUCCESS"
					break
				} else {
					fmt.Println("注册失败，5s后重新注册")
				}

				select {
				case <-s.regChan:
					return
				case <-time.After(5 * time.Second):
				}
			}
			isReReg = false
		}
		select {
		case sig := <-s.regChan:
			switch sig {
			case "stop", "unreg":
				return
			case "rereg":
				isReReg = true
				fmt.Println("重注册")
			}
		case <-time.After(time.Second):
			continue
		}
	}
}

// 对上级心跳检测
func (s *Superior) heartCheck() {
	defer func() {
		// heartCheck退出由heartBeatChan控制，退出时通知注册协程一起退出
		s.regChan <- "stop"
	}()

	//等待注册完成
	for s.regStatus != "SUCCESS" {
		fmt.Println("还没注册成功，等待注册成功才进行心跳检测")
		select {
		case _, ok := <-s.heartBeatChan:
			if !ok {
				return
			}
		case <-time.After(5 * time.Second):
		}
	}
	fmt.Println("注册成功，可以心跳检测了")

	time.Sleep(5 * time.Second)
	hearBeatInternal := 5

	for {
		// 向上级发心跳
		state := s.state
		if state == "offline" { // 上级下线
			fmt.Println("检测上级心跳失败")
			state = "offline"
		} else {
			fmt.Println("上级还活着")
		}
		if s.state == "offline" {
			fmt.Println("上级下线了，需要重新向上级注册")
			s.regChan <- "rereg"
		}

		select {
		case _, ok := <-s.heartBeatChan:
			if !ok {
				return
			}
		case <-time.After(time.Duration(hearBeatInternal) * time.Second):
		}
	}
}

// 模拟注册
func (u *userAgent) register() error {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(101)
	if randomNum%2 == 1 {
		return errors.New("register failed")
	}
	return nil
}

func (s *Superior) MakeSuperiorDown() {
	for {
		if Random() && s.state == "online" {
			s.state = "offline"
			fmt.Println("【上级下线】")
		}

		if Random() && s.state == "offline" {
			s.state = "online"
			fmt.Println("【上级上线】")
		}

		select {
		case _, ok := <-s.makeSuperiorDownChan:
			if !ok {
				return
			}
		case <-time.After(5 * time.Second):
		}
	}
}

func Random() bool {
	rand.Seed(time.Now().UnixNano())
	// 生成一个 0 到 n 的随机整数
	randomNum := rand.Intn(20)

	// 判断随机数是否为 0，以达到约 1/20 的概率
	return randomNum == 1
}
