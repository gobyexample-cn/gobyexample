# 当我们运行这个程序时，首先会看到阻塞式调用的输出，然后是两个协程的交替输出。
# 这种交替的情况表示 Go runtime 是以并发的方式运行协程的。
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# 接下来我们将学习协程的辅助特性：通道（channels）。
