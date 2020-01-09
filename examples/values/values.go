// Go 拥有多种值类型，包括字符串、整型、浮点型、布尔型等。
// 下面是一些基础的例子。

package main

import "fmt"

func main() {

	// 字符串可以通过 `+` 连接。
	fmt.Println("go" + "lang")

	// 整数和浮点数
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// 布尔型，以及常见的布尔操作。
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
