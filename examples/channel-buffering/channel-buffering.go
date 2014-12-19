// 默认通道是 _无缓冲_ 的，这意味着只有在对应的接收（`<- chan`）
// 通道准备好接收时，才允许进行发送（`chan <-`）。_可缓存通道_
// 允许在没有对应接收方的情况下，缓存限定数量的值。

package main

import "fmt"

func main() {

    // 这里我们 `make` 了一个通道，最多允许缓存 2 个值。
    messages := make(chan string, 2)

    // 因为这个通道是有缓冲区的，即使没有一个对应的并发接收
    // 方，我们仍然可以发送这些值。
    messages <- "buffered"
    messages <- "channel"

    // 然后我们可以像前面一样接收这两个值。
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
