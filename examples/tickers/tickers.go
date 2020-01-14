// [定时器](timers) 是当你想要在未来某一刻执行一次时使用的
// - _打点器_ 则是为你想要以固定的时间间隔重复执行而准备的。
// 这里是一个打点器的例子，它将定时的执行，直到我们将它停止。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 打点器和定时器的机制有点相似：使用一个通道来发送数据。
	// 这里我们使用通道内建的 `select`，等待每 500ms 到达一次的值。
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// 打点器可以和定时器一样被停止。
	// 打点器一旦停止，将不能再从它的通道中接收到值。
	// 我们将在运行 1600ms 后停止这个打点器。
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
