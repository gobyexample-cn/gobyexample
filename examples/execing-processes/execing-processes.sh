# 当我们运行程序时，它会替换为 `ls`。
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# 注意 Go 没有提供 Unix 经典的  `fork` 函数。
# 一般来说，这没有问题，因为启动协程、生成进程和执行进程，
# 已经涵盖了 fork 的大多数使用场景。
