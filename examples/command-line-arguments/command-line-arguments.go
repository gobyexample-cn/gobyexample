// [_命令行参数_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// 是一个指定程序运行参数的常见方式。例如，`go run hello.go`，
// 程序 `go` 使用 `run` 和 `hello.go` 两个参数。

package main

import "os"
import "fmt"

func main() {

	// `os.Args` 提供原始命令行参数访问功能。注意，切片中
	// 的第一个参数是该程序的路径，并且 `os.Args[1:]`保存
	// 所有程序的的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 你可以使用标准的索引位置取得单个参数的值。
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
