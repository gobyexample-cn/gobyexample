// 有时候，我们希望 Go 可以智能的处理 [Unix 信号](http://en.wikipedia.org/wiki/Unix_signal)。
// 例如，我们希望当服务器接收到一个 `SIGTERM` 信号时，能够优雅退出，
// 或者一个命令行工具在接收到一个 `SIGINT` 信号时停止处理输入信息。
// 我们这里讲的就是在 Go 中如何使用通道来处理信号。

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Go 通过向一个通道发送 `os.Signal` 值来发送信号通知。
	// 我们将创建一个通道来接收这些通知（同时再创建一个在程序结束时发送通知的通道）。
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` 注册给定的通道，用于接收特定信号。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 这个协程执行一个阻塞的信号接收操作。
	// 当它接收到一个值时，它将打印这个值，然后通知程序可以退出。
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// 程序将在这里进行等待，直到它得到了期望的信号
	// （也就是上面的协程发送的 `done` 值），然后退出。
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
