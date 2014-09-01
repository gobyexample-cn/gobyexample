# 注意，在使用 `fmt.Println` 来打印数组的时候，会使用
# `[v1 v2 v3 ...]` 的格式显示
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]

# 在典型的 Go 程序中，相对于数组而言，_slice_ 使用
# 的更多。我们将在后面讨论 slice。
