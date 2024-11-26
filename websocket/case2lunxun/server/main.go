package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	inventory int
	mu        sync.Mutex
	clients   = make(map[http.ResponseWriter]bool)
	clientsMu sync.Mutex
)

func updateInventory() {
	for {
		// 模拟库存变化
		mu.Lock()
		inventory++
		mu.Unlock()
		fmt.Printf("Inventory updated to %d\n", inventory)

		// 广播给所有客户端（长轮询中实际上是通过响应新请求来实现的）
		// 注意：这里并没有真正的“广播”机制，因为长轮询是每个客户端自己发起请求的

		// 等待一段时间再变化库存
		time.Sleep(2 * time.Second)
	}
}

func handleLongPoll(w http.ResponseWriter, r *http.Request) {
	// 这里应该有一个更复杂的逻辑来处理客户端断开连接的情况
	// 以及如何管理多个客户端的长轮询请求

	// 保持连接打开，直到有新的库存数据
	for {
		// 检查库存是否有更新（在实际应用中，您可能需要一个更高效的机制来通知客户端）
		mu.Lock()
		currentInventory := inventory
		mu.Unlock()

		// 假设我们有一种机制来检测“新”的库存（比如时间戳或版本号）
		// 这里我们简单地每次循环都发送最新的库存值作为示例

		// 设置响应头（这里应该还有Content-Type等其他头）
		w.Header().Set("Content-Type", "application/json")

		// 发送库存数据给客户端
		fmt.Fprintf(w, `{"inventory": %d}`, currentInventory)

		// 在长轮询中，我们通常在这里关闭连接，并让客户端重新发起请求
		// 但为了模拟持续连接的效果（尽管这不是真正的持续连接），
		// 我们可以让服务器等待一段时间再发送响应，模拟“等待新数据”的过程
		// 注意：在实际的长轮询实现中，您不应该这样做，而是应该让服务器
		// 在有新数据时立即响应，或者设置一个合理的超时时间。
		time.Sleep(1 * time.Second) // 模拟等待新数据的时间

		// 在真实场景中，您应该检查w.Write()的返回值以及w.Close()来确保连接正确关闭
		// 并且处理可能的错误。但由于这个示例的简化性质，我们省略了这些步骤。

		// 注意：由于HTTP连接的短暂性，上面的代码在每次循环结束时实际上都会关闭连接
		// 并导致客户端收到响应后重新发起请求。这就是长轮询的基本工作原理。
		// 如果您想要保持一个持久的连接（尽管不是真正的实时连接），您应该使用WebSocket等技术。

		// 在这个简化的示例中，我们模拟了长轮询的效果，但实际上并没有保持连接打开。
		// 为了保持连接打开，您需要实现一个更复杂的机制，比如使用goroutines和channels来管理客户端连接和消息传递。

		// 由于这个示例的局限性，我们将在每次循环结束时“断开”连接（即发送响应并等待客户端重新连接）
		// 在真实的应用中，您应该使用更合适的技术（如WebSocket）来实现持久连接和实时通信。

		// 注意：下面的break语句实际上会终止这个for循环，导致每次请求只发送一次响应。
		// 在真实的长轮询实现中，您应该移除这个break语句，并让循环继续运行，直到有新的数据可供发送或超时发生。
		// 但由于这个示例的简化性质，我们保留了break语句来模拟单次请求-响应周期。
		break // 移除这个break语句以实现真正的长轮询。
	}

	// 注意：在真实的应用中，您应该在这里添加适当的错误处理和日志记录。
}

func main() {
	http.HandleFunc("/poll", handleLongPoll)

	// 启动一个goroutine来模拟库存变化
	go updateInventory()

	// 启动HTTP服务器
	http.ListenAndServe(":8080", nil)
}
