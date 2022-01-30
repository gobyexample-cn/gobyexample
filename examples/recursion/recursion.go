// Go 支持 <a href="http://zh.wikipedia.org/wiki/%E9%80%92%E5%BD%92_(%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%A7%91%E5%AD%A6)"><em>递归</em></a>。
// 这里是一个经典的示例。

package main

import "fmt"

// `fact` 函数在到达 `fact(0)` 前一直调用自身。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// 闭包也可以是递归的，但这要求在定义闭包之前用类型化的 `var` 显式声明闭包。
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)

		// 由于 `fib` 之前在 `main` 中声明过，
		// 因此 Go 知道在这里用 `fib` 调用哪个函数。
	}

	fmt.Println(fib(7))
}
