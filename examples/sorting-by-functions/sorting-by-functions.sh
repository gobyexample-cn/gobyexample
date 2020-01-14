# 运行这个程序，和预期的一样，
# 显示了一个按照字符串长度排序的列表。
$ go run sorting-by-functions.go 
[kiwi peach banana]

# 类似的，参照这个例子，创建一个自定义类型，
# 为它实现 `Interface` 接口的三个方法，
# 然后对这个自定义类型的切片调用 `sort.Sort` 方法，
# 我们就可以通过任意函数对 Go 切片进行排序了。
