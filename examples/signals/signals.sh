# 当我们运行这个程序时，它将一直等待一个信号。
# 通过 `ctrl-C`（终端显示为 `^C`），我们可以发送一个 `SIGINT` 信号，
# 这会使程序打印 `interrupt` 然后退出。
$ go run signals.go
awaiting signal
^C
interrupt
exiting
