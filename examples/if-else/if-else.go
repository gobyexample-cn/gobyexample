// `if` 和 `else` 分支结构在 Go 中非常直接。

package main

import "fmt"

func main() {

	// 这里是一个基本的例子。
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 你可以不要 `else` 只用 `if` 语句。
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 在条件语句之前可以有一个声明语句；在这里声明的变量可以在这个语句所有的条件分支中使用。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// 注意，在 Go 中，条件语句的圆括号不是必需的，但是花括号是必需的。
