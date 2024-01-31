// `//go:embed` 是一个 [编译器指令](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives)，
// 它允许程序在构建时将任意文件和文件夹包含在 Go 二进制文件中。在这里阅读有关嵌入指令的更多信息：[这里](https://pkg.go.dev/embed)。
package main

// 导入 `embed` 包；如果您不使用该包中的任何导出标识符，可以使用 `_ "embed"` 进行空白导入。
import (
	"embed"
)

// `embed` 指令接受相对于包含 Go 源文件的目录的路径。
// 该指令将文件的内容嵌入到紧随其后的 `string` 变量中。
//
//go:embed folder/single_file.txt
var fileString string

// 将文件的内容嵌入到一个 `[]byte` 中。
//
//go:embed folder/single_file.txt
var fileByte []byte

// 我们还可以使用通配符嵌入多个文件甚至文件夹。
// 这将使用 [embed.FS 类型](https://pkg.go.dev/embed#FS)的变量，
// 该类型实现了一个简单的虚拟文件系统。
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// 打印出 `single_file.txt` 的内容。
	print(fileString)
	print(string(fileByte))

	// 从嵌入的文件夹中检索一些文件。
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
