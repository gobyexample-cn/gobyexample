# 运行程序将会引起 panic，输出一个错误消息和 Go 运行时
# 栈信息，并且返回一个非零的状态码。
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# 注意，不像有些语言中处理多个错误那样，在 Go 中习惯使用错
# 误码返回任意可能的值。
