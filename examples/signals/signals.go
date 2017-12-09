// 有时候，我们希望 Go 能智能的处理 <a href="http://zh.wikipedia.org/wiki/%E4%BF%A1%E5%8F%B7_(%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%A7%91%E5%AD%A6)">Unix 信号</a>。
// 例如，我们希望当服务器接收到一个 `SIGTERM` 信号时能够
// 自动关机，或者一个命令行工具在接收到一个 `SIGINT` 信号
// 时停止处理输入信息。这里讲的就就是在 Go 中如何通过通道
// 来处理信号。

package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

    // Go 通过向一个通道发送 `os.Signal` 值来进行信号通知。我们
    // 将创建一个通道来接收这些通知（同时还创建一个用于在程序可
    // 以结束时进行通知的通道）。
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    // `signal.Notify` 注册这个给定的通道用于接收特定信号。
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // 这个 Go 协程执行一个阻塞的信号接收操作。当它得到一个
    // 值时，它将打印这个值，然后通知程序可以退出。
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    // 程序将在这里进行等待，直到它得到了期望的信号（也就
    // 是上面的 Go 协程发送的 `done` 值）然后退出。
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
