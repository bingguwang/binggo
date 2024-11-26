package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

/*
*
编写一个客户端，模拟下访问接口
*/
func main() {
	for {
		wg := sync.WaitGroup{}
		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				resp, err := http.Get("http://192.168.0.66:20001/hello")
				if err != nil {
					log.Println(err)
					return
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Println(err)
					return
				}

				fmt.Println(string(body))
			}(i)
		}

		wg.Wait()
		time.Sleep(5 * time.Second)
	}

}
