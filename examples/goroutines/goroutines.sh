# 当我们运行这个程序时，将首先看到阻塞式调用的输出，然后是
# 两个 Go 协程的交替输出。这种交替的情况表示 Go 运行时是以
# 异步的方式运行协程的。
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
<enter>
done

# 接下来我们将学习在并发的 Go 程序中的 Go 协程的辅助
# 特性：通道。
