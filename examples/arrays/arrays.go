// 在 Go 中，_数组_ 是一个固定长度的数列。

package main

import "fmt"

func main() {

    // 这里我们创建了一个数组 `a` 来存放刚好 5 个 `int`。
    // 元素的类型和长度都是数组类型的一部分。数组默认是
    // 零值的，对于 `int` 数组来说也就是 `0`。
    var a [5]int
    fmt.Println("emp:", a)

    // 我们可以使用 `array[index] = value` 语法来设置数组
    // 指定位置的值，或者用 `array[index]` 得到值。
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // 使用内置函数 `len` 返回数组的长度
    fmt.Println("len:", len(a))

    // 使用这个语法在一行内初始化一个数组
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // 数组的存储类型是单一的，但是你可以组合这些数据
    // 来构造多维的数据结构。
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
