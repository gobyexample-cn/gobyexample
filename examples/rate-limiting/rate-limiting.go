// <em>[速率限制](http://en.wikipedia.org/wiki/Rate_limiting)</em>
// 是控制服务资源利用和质量的重要机制。
// 基于协程、通道和[打点器](tickers)，Go 优雅的支持速率限制。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 首先，我们将看一个基本的速率限制。
	// 假设我们想限制对收到请求的处理，我们可以通过一个渠道处理这些请求。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// `limiter` 通道每 200ms 接收一个值。
	// 这是我们任务速率限制的调度器。
	limiter := time.Tick(200 * time.Millisecond)

	// 通过在每次请求前阻塞 `limiter` 通道的一个接收，
	// 可以将频率限制为，每 200ms 执行一次请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们想临时进行速率限制，并且不影响整体的速率控制，
	// 我们可以通过[通道缓冲](channel-buffering.html)来实现。
	// 这个 `burstyLimiter` 通道用来进行 3 次临时的脉冲型速率限制。
	burstyLimiter := make(chan time.Time, 3)

	// 想将通道填充需要临时改变3次的值，做好准备。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每 200 ms 我们将添加一个新的值到 `burstyLimiter`中，
	// 直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 现在模拟超过 5 个的接入请求。它们中刚开始的 3 个将
	// 由于受 `burstyLimiter` 的“脉冲”影响。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
