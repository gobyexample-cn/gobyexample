# 运行程序将会导致 panic：
# 输出一个错误消息和协程追踪信息，并以非零的状态退出程序。

# 当 `main` 中触发第一个 panic 时，程序就会退出而不会执行代码的其余部分。
# 如果你想看到程序尝试创建 /tmp/file 文件，请注释掉第一个panic。
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# 注意，与某些使用 exception 处理错误的语言不同，
# 在 Go 中，通常会尽可能的使用返回值来标示错误。
