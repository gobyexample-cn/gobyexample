# 我们首先接收到值 `"one"`，然后就是预料中的 `"two"`
# 了。
$ time go run select.go 
received one
received two

# 注意从第一次和第二次 `Sleeps` 并发执行，总共仅运行了
# 两秒左右。
real	0m2.245s
