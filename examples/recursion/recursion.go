// Go 支持 <a href="http://zh.wikipedia.org/wiki/%E9%80%92%E5%BD%92_(%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%A7%91%E5%AD%A6)"><em>递归</em></a>。
// 这里是一个经典的阶乘示例。

package main

import "fmt"

// `face` 函数在到达 `face(0)` 前一直调用自身。
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))
}
