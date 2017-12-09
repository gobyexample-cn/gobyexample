// 我们常常需要在后面一个时刻运行 Go 代码，或者在某段时间
// 间隔内重复运行。Go 的内置 _定时器_ 和 _打点器_ 特性让这
// 些很容易实现。我们将先学习定时器，然后再学习[打点器](../tickers/)。

package main

import "time"
import "fmt"

func main() {

    // 定时器表示在未来某一时刻的独立事件。你告诉定时器
    // 需要等待的时间，然后它将提供一个用于通知的通道。
    // 这里的定时器将等待 2 秒。
    timer1 := time.NewTimer(time.Second * 2)

    // `<-timer1.C` 直到这个定时器的通道 `C` 明确的发送了
    // 定时器失效的值之前，将一直阻塞。
    <-timer1.C
    fmt.Println("Timer 1 expired")

    // 如果你需要的仅仅是单纯的等待，你需要使用 `time.Sleep`。
    // 定时器是有用原因之一就是你可以在定时器失效之前，取消这个
    // 定时器。这是一个例子
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
