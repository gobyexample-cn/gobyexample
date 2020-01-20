# 要实验命令行参数，最好先使用 `go build` 编译一个可执行二进制文件
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]
[a b c d]
c

# 下面我们要看看更高级的使用标记的命令行处理方法。
