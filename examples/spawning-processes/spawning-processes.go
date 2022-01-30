// 有时，我们的 Go 程序需要生成其他的、非 Go 的进程。

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// 我们将从一个简单的命令开始，没有参数或者输入，仅打印一些信息到标准输出流。
	// `exec.Command` 可以帮助我们创建一个对象，来表示这个外部进程。
	dateCmd := exec.Command("date")

	// `.Output` 是另一个帮助函数，常用于处理运行命令、等待命令完成并收集其输出。
	// 如果没有错误，`dateOut` 将保存带有日期信息的字节。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 下面我们将看看一个稍复杂的例子，
	// 我们将从外部进程的 `stdin` 输入数据并从 `stdout` 收集结果。
	grepCmd := exec.Command("grep", "hello")

	// 这里我们明确的获取输入/输出管道，运行这个进程，
	// 写入一些输入数据、读取输出结果，最后等待程序运行结束。
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// 上面的例子中，我们忽略了错误检测，
	// 当然，你也可以使用常见的 `if err != nil` 方式来进行错误检查。
	// 我们只收集了 `StdoutPipe` 的结果，
	// 但是你可以使用相同的方法收集 `StderrPipe` 的结果。
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 注意，在生成命令时，我们需要提供一个明确描述命令和参数的数组，而不能只传递一个命令行字符串。
	// 如果你想使用一个字符串生成一个完整的命令，那么你可以使用 `bash` 命令的 `-c` 选项：
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
