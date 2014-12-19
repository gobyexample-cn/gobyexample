// Go 支持通过 <a href="http://zh.wikipedia.org/wiki/%E9%97%AD%E5%8C%85_(%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%A7%91%E5%AD%A6)"><em>闭包</em></a>
// 来使用 [_匿名函数_](http://zh.wikipedia.org/wiki/%E5%8C%BF%E5%90%8D%E5%87%BD%E6%95%B0)。
// 匿名函数在你想定义一个不需要命名的内联函数时是很实用的。

package main

import "fmt"

// 这个 `intSeq` 函数返回另一个在 `intSeq` 函数体内定义的
// 匿名函数。这个返回的函数使用闭包的方式 _隐藏_ 变量 `i`。
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {

    // 我们调用 `intSeq` 函数，将返回值（也是一个函数）赋给
    // `nextInt`。这个函数的值包含了自己的值 `i`，这样在每
    // 次调用 `nextInt` 是都会更新 `i` 的值。
    nextInt := intSeq()

    // 通过多次调用 `nextInt` 来看看闭包的效果。
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // 为了确认这个状态对于这个特定的函数是唯一的，我们
    // 重新创建并测试一下。
    newInts := intSeq()
    fmt.Println(newInts())
}
