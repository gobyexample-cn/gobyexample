// Go 的 `sort` 包实现了内建及用户自定义数据类型的排序功能。
// 我们先来看看内建数据类型的排序。

package main

import (
	"fmt"
	"sort"
)

func main() {

	// 排序方法是针对内置数据类型的；
	// 这是一个字符串排序的例子。
	// 注意，它是原地排序的，所以他会直接改变给定的切片，而不是返回一个新切片。
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 一个 `int` 排序的例子。
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// 我们也可以使用 `sort` 来检查一个切片是否为有序的。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
