# 跟预期的一样，我们首先接收到值 `"one"`，然后是 `"two"`。
$ time go run select.go 
received one
received two

# 注意，程序总共仅运行了两秒左右。因为 1 秒 和 2 秒的 `Sleeps` 是并发执行的，
real	0m2.245s
