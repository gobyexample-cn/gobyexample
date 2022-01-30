// Go 在传统的 `printf` 中对字符串格式化提供了优异的支持。
// 这儿有一些基本的字符串格式化的任务的例子。

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go 提供了一些用于格式化常规值的打印“动词”。
	// 例如，这样打印 `point` 结构体的实例。
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// 如果值是一个结构体，`%+v` 的格式化输出内容将包括结构体的字段名。
	fmt.Printf("struct2: %+v\n", p)

	// `%#v` 根据 Go 语法输出值，即会产生该值的源码片段。
	fmt.Printf("struct3: %#v\n", p)

	// 需要打印值的类型，使用 `%T`。
	fmt.Printf("type: %T\n", p)

	// 格式化布尔值很简单。
	fmt.Printf("bool: %t\n", true)

	// 格式化整型数有多种方式，使用 `%d` 进行标准的十进制格式化。
	fmt.Printf("int: %d\n", 123)

	// 这个输出二进制表示形式。
	fmt.Printf("bin: %b\n", 14)

	// 输出给定整数的对应字符。
	fmt.Printf("char: %c\n", 33)

	// `%x` 提供了十六进制编码。
	fmt.Printf("hex: %x\n", 456)

	// 同样的，也为浮点型提供了多种格式化选项。
	// 使用 `%f` 进行最基本的十进制格式化。
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` 和 `%E` 将浮点型格式化为（稍微有一点不同的）科学记数法表示形式。
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// 使用 `%s` 进行基本的字符串输出。
	fmt.Printf("str1: %s\n", "\"string\"")

	// 像 Go 源代码中那样带有双引号的输出，使用 `%q`。
	fmt.Printf("str2: %q\n", "\"string\"")

	// 和上面的整型数一样，`%x` 输出使用 base-16 编码的字符串，
	// 每个字节使用 2 个字符表示。
	fmt.Printf("str3: %x\n", "hex this")

	// 要输出一个指针的值，使用 `%p`。
	fmt.Printf("pointer: %p\n", &p)

	// 格式化数字时，您经常会希望控制输出结果的宽度和精度。
	// 要指定整数的宽度，请在动词 "%" 之后使用数字。
	// 默认情况下，结果会右对齐并用空格填充。
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// 你也可以指定浮点型的输出宽度，同时也可以通过 `宽度.精度` 的语法来指定输出的精度。
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// 要左对齐，使用 `-` 标志。
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。
	// 这是基本的宽度右对齐方法。
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// 要左对齐，和数字一样，使用 `-` 标志。
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// 到目前为止，我们已经看过 `Printf` 了，
	// 它通过 `os.Stdout` 输出格式化的字符串。
	// `Sprintf` 则格式化并返回一个字符串而没有任何输出。
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// 你可以使用 `Fprintf` 来格式化并输出到 `io.Writers` 而不是 `os.Stdout`。
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
