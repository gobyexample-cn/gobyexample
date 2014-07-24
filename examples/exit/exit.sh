# 如果你使用 `go run` 来运行 `exit.go`，那么退出状态将会被 `go`
# 捕获并打印。
$ go run exit.go
exit status 3

# 使用编译并执行一个二进制文件的方式，你可以在终端中查看退出状态。
$ go build exit.go
$ ./exit
$ echo $?
3

# 注意我们程序中的 `!` 永远不会被打印出来。
