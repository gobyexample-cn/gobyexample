// Go标准库提供了直观的工具用于从Go程序输出日志
// 使用 [log](https://pkg.go.dev/log) 包进行自由格式输出
// 使用 [log/slog](https://pkg.go.dev/log/slog) 包进行结构化输出。

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// 只需调用 `log` 包中的 `Println` 等函数即可使用 _标准_ logger。
	// 它已经预先配置为将日志输出到 `os.Stderr`。
	// 像 `Fatal*` 或 `Panic*` 这样的附加方法将在记录日志后退出程序。
	log.Println("standard logger")

	// 日志记录器可以使用 _flags_ 进行配置，以设置它们的输出格式。
	// 默认情况下，标准记录器已设置了 `log.Ldate` 和 `log.Ltime` 标志，
	// 并将它们收集在 `log.LstdFlags` 中。
	// 我们可以更改其标志以发出微秒精度的时间，例如：
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// 它还支持发出调用 log` 函数的文件名和行号。
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// 可能会有用创建一个自定义记录器并在各处传递它。
	// 创建新记录器时，我们可以设置一个 _前缀_ 来区分其输出和其他日志记录器。
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// 我们可以使用 `SetPrefix` 方法在现有的记录器（包括标准记录器）上设置前缀。
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// 日志记录器可以具有自定义的输出目标；任何 `io.Writer` 都可以使用。
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// 这个调用将日志输出写入到 `buf` 中.
	buflog.Println("hello")

	// 这将实际上显示在标准输出上。
	fmt.Print("from buflog:", buf.String())

	// `slog 包提供了 _结构化_ 的日志输出。例如，以 JSON 格式记录日志非常直接。
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// 除了 `msg` 之外，`slog` 输出还可以包含任意数量的键值对。
	myslog.Info("hello again", "key", "val", "age", 25)
}
