// _关闭_ 一个通道意味着不能再向这个通道发送值了。这个可以
// 用来给这个通道的接收方传达工作已将完成。

package main

import "fmt"

// 在这个例子中，我们将使用一个 `jobs` 通道来传递 `main()` Go
// 协程中的任务执行结束信息到一个工作 Go 协程中。当我们没有多余的
// 任务给 worker 时，我们将 `close` 这个 `jobs` 通道。
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // 这是工作 Go 协程。使用 `j, more := <- jobs` 从 `jobs`
    // 反复的进行接收。在接收的这的这个特殊的二值形式的值中，
    // 如果 `jobs` 已经关闭了，并且通道中所有的值都已经被接收
    // 到了，那么 `more` 的值将设置为 `false`。当我们做完所有
    // 的任务时，将使用这个在 `done` 中去进行通知。
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // 这里使用 `jobs` 发送 3 个任务到工作函数中，然后
    // 关闭 `jobs`。
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // 我们使用前面学到的[通道同步](channel-synchronization.html)
    // 方法等待任务结束。
    <-done
}
