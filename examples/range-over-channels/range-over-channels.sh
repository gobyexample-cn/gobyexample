$ go run range-over-channels.go
one
two

# 这个例子也让我们看到，一个非空的通道也是可以关闭的，
# 并且，通道中剩下的值仍然可以被接收到。
