# 后台运行服务器。
$ go run context-in-http-servers.go &

# 模拟客户端发出 `/hello` 请求，
# 在服务端开始处理后，按下 Ctrl+C 以发出取消信号。
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
