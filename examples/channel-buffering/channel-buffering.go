// 默认情况下，通道是 _无缓冲_ 的，这意味着只有对应的接收（`<- chan`）
// 通道准备好接收时，才允许进行发送（`chan <-`）。
// _可缓存通道_ 允许在没有对应接收者的情况下，缓存一定数量的值。

package main

import "fmt"

func main() {

	// 这里我们 `make` 了一个字符串通道，最多允许缓存 2 个值。
	messages := make(chan string, 2)

	// 由于此通道是带缓冲的，
	// 因此我们可以将这些值发送到通道中而不需要并发接收。
	messages <- "buffered"
	messages <- "channel"

	// 然后我们可以正常接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
