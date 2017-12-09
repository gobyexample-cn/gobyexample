// _通道_ 是连接多个 Go 协程的管道。你可以从一个 Go 协程
// 将值发送到通道，然后在别的 Go 协程中接收。

package main

import "fmt"

func main() {

    // 使用 `make(chan val-type)` 创建一个新的通道。
    // 通道类型就是他们需要传递值的类型。
    messages := make(chan string)

    // 使用 `channel <-` 语法 _发送_ 一个新的值到通道中。这里
    // 我们在一个新的 Go 协程中发送 `"ping"` 到上面创建的
    // `messages` 通道中。
    go func() { messages <- "ping" }()

    // 使用 `<-channel` 语法从通道中 _接收_ 一个值。这里
    // 将接收我们在上面发送的 `"ping"` 消息并打印出来。
    msg := <-messages
    fmt.Println(msg)
}
