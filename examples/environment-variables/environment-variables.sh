# 运行这个程序，显示我们在程序中设置的 `FOO` 的值，
# 然而没有设置的 `BAR` 是空的。
$ go run environment-variables.go
FOO: 1
BAR: 

# 键的列表是由你的电脑情况而定的。
TERM_PROGRAM
PATH
SHELL
...

# 如果我们在运行前设置了 `BAR` 的值，那么运行程序将会获取到这个值。
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
