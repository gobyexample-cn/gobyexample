// 第一个定时器将在程序开始后 ~2s 失效，但是第二个在它
// 没失效之前就停止了。
$ go run timers.go
Timer 1 expired
Timer 2 stopped
