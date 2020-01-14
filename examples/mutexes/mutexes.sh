# 运行这个程序，显示我们进行了大约 90,000 次 `mutex` 同步的 `state` 操作。
$ go run mutexes.go
readOps: 83285
writeOps: 8320
state: map[1:97 4:53 0:33 2:15 3:2]

# 接下来我们将看一下，只使用协程和通道，
# 如何实现相同的任务状态管理。
