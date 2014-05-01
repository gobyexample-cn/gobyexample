# 当我们运行程序师，它会替换为 `ls`。
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# 注意 Go 并不提供一个经典的 Unix `fork` 函数。通常这不
# 是个问题，因为运行 Go 协程，生成进程和执行进程覆盖了
# fork 的大多数使用用场景。
