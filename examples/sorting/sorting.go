// Go 的 `sort` 包实现了内置和用户自定义数据类型的排序
// 功能。我们首先关注内置数据类型的排序。

package main

import "fmt"
import "sort"

func main() {

    // 排序方法是正对内置数据类型的；这里是一个字符串的例子。
    // 注意排序是原地更新的，所以他会改变给定的序列并且不返回
    // 一个新值。
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    // 一个 `int` 排序的例子。
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    // 我们也可以使用 `sort` 来检查一个序列是不是已经
    // 是排好序的。
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
