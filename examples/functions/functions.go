// _函数_ 是 Go 的中心。我们将通过一些不同的例子来
// 进行学习。

package main

import "fmt"

// 这里是一个函数，接受两个 `int` 并且以 `int` 返回它
// 们的和
func plus(a int, b int) int {

    // Go 需要明确的返回值，例如，它不会自动返回最
    // 后一个表达式的值
    return a + b
}

func main() {

    // 正如你期望的那样，通过 `name(args)` 来调用一
    // 个函数，
    res := plus(1, 2)
    fmt.Println("1+2 =", res)
}

// todo: coalesced parameter types
