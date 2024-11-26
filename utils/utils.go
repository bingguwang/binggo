package utils

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
)

// GetOutBoundIP 获取到对外的ip
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "6.6.6.6:6666")
	if err != nil {
		log.Println("GetOutBoundIP failed:", err.Error())
		return
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), err
}

/**
 * @Description: 获取节点UUID（非虚拟机UUID）
 * @return string: 空表示获取失败
 */
func GetNodeUUID() (string, error) {
	switch OS := runtime.GOOS; OS {
	case "linux":
		// 通过文件存储UUID
		uuidFile := "/conf/uuid"
		if err := os.MkdirAll("/conf", 0766); err != nil {
			return "", err
		}

		if _, statErr := os.Stat(uuidFile); os.IsNotExist(statErr) {
			return func() (string, error) {
				f, err := os.Create(uuidFile)
				if err != nil {
					return "", err
				}
				defer f.Close()

				uuid, err := GenerateUUID()
				if err != nil {
					return "", err
				}

				if _, err = f.WriteString(uuid); err != nil {
					return "", err
				}

				return uuid, nil
			}()
		} else {
			return func() (string, error) {
				f, err := os.OpenFile(uuidFile, os.O_RDONLY, 0666)
				if err != nil {
					return "", nil
				}
				defer f.Close()

				if buf, err := io.ReadAll(f); err != nil {
					return "", nil
				} else {
					return string(buf), nil
				}
			}()
		}

	case "windows":
		//out, err := Exec("wmic", "csproduct", "list", "full")
		// 仅做测试用
		return GenerateUUID()
	default:
		// 不支持其余操作系统
		fmt.Println(OS)
		return "", errors.New("unsupported OS type")
	}
}
