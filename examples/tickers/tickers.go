// [定时器](../timers/) 是当你想要在未来某一刻执行一次时
// 使用的 - _打点器_ 则是当你想要在固定的时间间隔重复执行
// 准备的。这里是一个打点器的例子，它将定时的执行，直到我
// 们将它停止。

package main

import "time"
import "fmt"

func main() {

	// 打点器和定时器的机制有点相似：一个通道用来发送数据。
	// 这里我们在这个通道上使用内置的 `range` 来迭代值每隔
	// 500ms 发送一次的值。
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// 打点器可以和定时器一样被停止。一旦一个打点停止了，
	// 将不能再从它的通道中接收到值。我们将在运行后 1600ms
	// 停止这个打点器。
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
