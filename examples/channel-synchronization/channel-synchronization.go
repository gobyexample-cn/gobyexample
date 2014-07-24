// 我们可以使用通道来同步 Go 协程之间执行。这里是一个
// 使用阻塞的接受方式来等待一个 Go 协程的运行结束。

package main

import "fmt"
import "time"

// 这是我们将要在一个 Go 协程中运行的函数。`done` 通道
// 将被用于通知其他 Go 协程这个函数已经工作完毕。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知我们已经完工啦。
	done <- true
}

func main() {

	// 运行一个 worker Go协程，并把用来通知的协程传递过去。
	done := make(chan bool, 1)
	go worker(done)

	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	<-done
}
