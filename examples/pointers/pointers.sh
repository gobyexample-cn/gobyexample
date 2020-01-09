# `zeroval` 在 `main` 函数中不能改变 `i` 的值，
# 但是 `zeroptr` 可以，因为它有这个变量的内存地址的引用。
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
