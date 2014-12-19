// `for` 是 Go 中唯一的循环结构。这里有 `for` 循环
// 的三个基本使用方式。

package main

import "fmt"

func main() {

    // 最常用的方式，带单个循环条件。
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // 经典的初始化/条件/后续形式 `for` 循环。
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // 不带条件的 `for` 循环将一直执行，直到在循环体内使用
    // 了 `break` 或者 `return` 来跳出循环。
    for {
        fmt.Println("loop")
        break
    }
}
