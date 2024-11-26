package struct2Map

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestNamesss(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		generateOrgGbId("5121100174132", 1)
	}
	fmt.Println("len:", len(mp))
	fmt.Println("len:", mp[gbKey{platID: 1, gbID: "5121100174132000002"}])
	fmt.Println("len:", mp[gbKey{platID: 1, gbID: "5121100174132001900"}])
}

type gbKey struct {
	platID int64
	gbID   string
}

var mp = map[gbKey]string{}

func generateOrgGbId(base string, toPlatformId int64) string {
	count, step := 1, 10
	gbID := base + fmt.Sprintf("%06d", count)
	for count < 1000000 {
		// 是否冲突
		if _, ok := mp[gbKey{platID: toPlatformId, gbID: gbID}]; !ok {
			fmt.Println("生成的国标id是:", gbID)
			mp[gbKey{platID: toPlatformId, gbID: gbID}] = gbID
			return gbID
		}
		count += step
		if step <= 100 {
			// 会丢失步长内的id
			step += rand.Intn(9) + 1
		} else {
			step += rand.Intn(2) + 1
		}
		gbID = base + fmt.Sprintf("%06d", count)
	}
	fmt.Println("相机数超过最大值1000000")
	return ""
}

func TestLASDK(t *testing.T) {
	subsc()
}

func subsc() {
	vcsRedisClient := redis.NewClient(&redis.Options{
		Addr: "192.168.3.154:6379",
		DB:   1,
	})
	fmt.Println("连接redis ", vcsRedisClient.Ping(context.Background()).Err())
	ctx := context.Background()
	redisClient := vcsRedisClient
	pubSub := redisClient.Subscribe(ctx, "GB_CHANNEL")
	defer pubSub.Close()
	fmt.Println("pubSub----", pubSub.String())
	for {
		select {
		case v := <-pubSub.Channel():
			msg := v.String()
			split := strings.Split(strings.Trim(msg, " "), "|")
			if len(split) < 3 {
				fmt.Printf("频道[%s]传递的信息格式有错误:%s\n", "GB_CHANNEL", msg)
				continue
			}
			fmt.Println("收到GBChannel消息: ", msg)
			operate := split[2]
			fmt.Println(operate)
		}
	}
}

func TestKsd(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exitChan := make(chan string)
	defer close(exitChan)
	fmt.Println(ctx.Value(""))
	var wg sync.WaitGroup
	go func(ctx context.Context, exitChan chan string) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second):
				intn := rand.Intn(10)
				if intn%6 == 0 {
					exitChan <- fmt.Sprintf("退出")
					return
				} else {
					fmt.Println("online......")
				}
			}
		}
	}(ctx, exitChan)
LOOP:
	for {
		select {
		case reason := <-exitChan:
			fmt.Println("exit:", reason)
			cancel()
			break LOOP
		}
	}
	wg.Wait()
}

func TestNamess(t *testing.T) {
	ptzCmd := "A50F008100010036"
	fmt.Println("16位的字符串是十六进制数，每个十六进制数是一个字节，共八个字节，需要转为10进制的")
	var checkNum, presetId int64
	var ctrlCode []int64
	for i := 0; i < 8; i++ {
		fmt.Println("【解析 ", ptzCmd[i*2:i*2+2], "】")
		cmdCode, _ := strconv.ParseInt(ptzCmd[i*2:i*2+2], 16, 64)
		fmt.Println("转为10进制: ", cmdCode)
		if i < 7 {
			checkNum = checkNum + cmdCode
		}
		ctrlCode = append(ctrlCode, cmdCode)
	}
	fmt.Println("16位:", ctrlCode)
	fmt.Println("checkNum是前七个字节的和:", checkNum)

	// 8th byte checksum
	fmt.Println("校验一下第八个字节，第八个字节应该 == 前七个字节的和 % 256")
	if ctrlCode[7] != (checkNum & 0xff) {
		fmt.Println("checksum error")
		return
	}

	// # 1st byte, always A5
	fmt.Println("检查第一个字节，第一个字节都是A5")
	if ctrlCode[0] != 0xA5 {
		fmt.Println("First control byte is not A5")
		return
	}

	// TODO 3th byte: low 8 bit of address
	fmt.Println("第3个字节, 代表这地址的低八位")

	fmt.Println("重点关注的是第4字节，这个字节表示这指令码")
	// 4th byte
	var zoomout, zoomin, up, down, left, right, irisShrink, irisExpand, focusIn, focusOut int64 = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	var byte4 = ctrlCode[3]

	// preset cmd in A.3.4
	// set or move or delete preset position
	if byte4 == 0x81 || byte4 == 0x82 || byte4 == 0x83 {
		// 预置位指令见A3.4
		// 第4字节
		// 0x81 设置预置位
		// 0x82 调用预置位
		// 0x83 删除预置位
		if ctrlCode[4] != 0 {
			fmt.Println("Invalid preset cmd,the fifth byte is not 0")
			return
		}
		fmt.Println("预置位指令的第5字节都是0x00")
		if ctrlCode[5] == 0 {
			fmt.Println("Invalid preset cmd,the sixth byte is 0")
			return
		}
		presetId = ctrlCode[5]
		fmt.Println("presetId:", presetId)
	} else if func(num int64) bool {
		checkCodes := []int64{0x84, 0x85, 0x86, 0x87, 0x88, 0x00}
		for _, element := range checkCodes {
			if num == element {
				return true
			}
		}
		return false
	}(byte4) {
		// 巡航指令见A3.5
		// 0x84 加入巡航点
		// 0x85 删除巡航点
		// 0x86 设置巡航速度
		// 0x87 设置巡航停留时间
		// 0x88 开始巡航
		// 0x00 停止巡航
		tourNum := ctrlCode[4]
		fmt.Println("第5个字节表示巡航组号：", tourNum)
		presetNum := ctrlCode[5]
		fmt.Println("第6个字节表示预置位号：", presetNum)
		if byte4 == 0x84 {
			fmt.Println("加入巡航点")
		} else if byte4 == 0x85 {
			fmt.Println("删除巡航点")
			if ctrlCode[5] == 0 {
				fmt.Println("第六个节点为0代表没指定是删除巡航的哪个预置位，表示删除整条巡航")
			}
		} else if byte4 == 0x86 {
			fmt.Println("设置巡航速度")
			// NOTE: 这里不应该有preset_num参数，ctrl_code[5]应该表示speed的低8位
			// TODO: VCS不支持速度，故这里未实现
			//speed := ctrlCode[6] & 0xf0

		} else if byte4 == 0x87 {
			fmt.Println("设置巡航停留时间")
			timeStr := strconv.Itoa(int(ctrlCode[6]/0xf)) + strconv.FormatInt(int64(ctrlCode[5]), 16)
			stayTime, _ := strconv.ParseInt(timeStr, 16, 64)
			fmt.Println("stayTime:", stayTime)
		} else if byte4 == 0x88 {
			fmt.Println("开始巡航")
		} else {
			fmt.Println("停止巡航")
		}
	} else if byte4&0xc0 == 0 {
		// PTZ指令见A3.2
		if byte4&0x30 == 0x30 {
			fmt.Println("zoomout zoomin can not both be 1")
			return
		}
		if byte4&0x0c == 0x0c {
			fmt.Println("up down can not both be 1")
			return
		}
		if byte4&0x03 == 0x03 {
			fmt.Println("left right can not both be 1")
			return
		}

		if byte4&0x20 != 0 {
			zoomout = (ctrlCode[6] & 0xf0) >> 4
		} else if byte4&0x10 != 0 {
			zoomin = (ctrlCode[6] & 0xf0) >> 4
		}

		if byte4&0x08 != 0 {
			up = ctrlCode[5]
		} else if byte4&0x04 != 0 {
			down = ctrlCode[5]
		}

		if byte4&0x02 != 0 {
			left = ctrlCode[4]
		} else if byte4&0x01 != 0 {
			right = ctrlCode[4]
		}
		fmt.Println("zoomout: ", zoomout)
		fmt.Println("zoomin: ", zoomin)
		fmt.Println("up: ", up)
		fmt.Println("down: ", down)
		fmt.Println("left: ", left)
		fmt.Println("right: ", right)
		fmt.Println("irisShrink: ", irisShrink)
		fmt.Println("irisExpand: ", irisExpand)
		fmt.Println("focusIn: ", focusIn)
		fmt.Println("focusOut: ", focusOut)
	} else if byte4&0xc0 == 0x40 {
		// FI指令见A3.3
		if byte4&0x0c == 0x0c {
			fmt.Println("iris shrink and expand can not both be 1")
			return
		}
		if byte4&0x03 == 0x03 {
			fmt.Println("focus in and out can not both be 1")
			return
		}
		if byte4&0x08 != 0 {
			irisShrink = ctrlCode[5]
		} else if byte4&0x04 != 0 {
			irisExpand = ctrlCode[5]
		} else if byte4&0x02 != 0 {
			focusIn = ctrlCode[4]
		} else if byte4&0x01 != 0 {
			focusOut = ctrlCode[4]
		}
	} else {
		fmt.Println("Invalid Byte 4 in control code")
		return
	}
	// send command
	if byte4&0xc0 == 0 {
		// ptz command
	} else if byte4&0xc0 == 0x40 {
		// fi command
	} else if byte4 == 0x81 {
		// set preset position
		fmt.Println("设置预置位")
	} else if byte4 == 0x82 {
		// move preset position
		fmt.Println("调用预置位")
	} else if byte4 == 0x83 {
		// delete preset position
		fmt.Println("删除预置位")
	}

}

func TestName222(t *testing.T) {

}
