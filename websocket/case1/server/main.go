package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
	inventory int
)
var (
	valChan = make(chan int)
)

func main() {
	http.HandleFunc("/subscribe", handleSubscribe)
	go simulateInventoryChanges()
	http.ListenAndServe(":8080", nil)

}

func handleSubscribe(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer ws.Close()

	clientsMu.Lock()
	clients[ws] = true
	clientsMu.Unlock()
	for {
		// 这里可以添加逻辑来处理连接断开的情况
		// 比如使用 select 和 done 通道来监听连接关闭事件
		// 但为了简化示例，我们直接在这里循环

		// 发送当前库存信息给客户端（在实际应用中，你可能只想在库存变化时发送）

		select {
		case <-valChan:
			for cli, b := range clients {
				if b {
					err = cli.WriteJSON(map[string]int{"inventory": inventory})
					if err != nil {
						break
					}
				}
			}
		}
		// 等待一段时间再发送（这里只是为了模拟实时更新，实际应用中你应该在库存变化时发送）
		//time.Sleep(1 * time.Second)
	}

	clientsMu.Lock()
	delete(clients, ws)
	clientsMu.Unlock()
}

func simulateInventoryChanges() {
	for {
		// 模拟库存变化（这里使用简单的递增和随机时间间隔）
		clientsMu.Lock()
		numClients := len(clients)
		clientsMu.Unlock()

		inventory++
		fmt.Printf("Inventory changed to %d (clients: %d)\n", inventory, numClients)

		// 广播库存变化给所有客户端（这里为了简化，没有实现真正的广播逻辑）
		// 你需要遍历clients map，并对每个客户端发送消息
		// 注意：在并发环境中，你需要确保对clients的访问是线程安全的

		// 为了简化示例，我们只等待一段时间再变化库存
		value := rand.Intn(5)
		time.Sleep(time.Duration(value) * time.Second)
		fmt.Println("写入通道 ", inventory, " time:", time.Now().String())
		valChan <- inventory
	}
}
