# 运行这个程序，首先显示运行超时的操作，然后是成功接收的。
$ go run timeouts.go 
timeout 1
result 2

# 使用这个 `select` 超时方式，需要使用通道传递结果。这对于
# 一般情况是个好的方式，因为其他重要的 Go 特性是基于通道和
# `select` 的。接下来我们就要看到两个例子：timer 和 ticker。
