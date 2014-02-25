# 注意，slice 和数组不同，虽然它们通过 `fmt.Println` 输出
# 差不多。
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# 看看这个由 Go 团队撰写的一篇[很棒的博文](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html)，
# 获得更多关于 Go 中 slice 的设计和实现细节。

# 现在，我们已经看过了数组和 slice，接下来我们将看看 
# Go 中的另一个关键的内建数据类型：map。