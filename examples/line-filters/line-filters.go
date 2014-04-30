// 一个_行过滤器_ 在读取标准输入流的输入，处理该输入，然后
// 将得到一些的结果输出到标准输出的程序中是常见的一个功能。
// `grep` 和 `sed` 是常见的行过滤器。

// 这里是一个使用 Go 编写的行过滤器示例，它将所有的输入文字
// 转化为大写的版本。你可以使用这个模式来写一个你自己的 Go
// 行过滤器。
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    // 对 `os.Stdin` 使用一个带缓冲的 scanner，让我们可以
    // 直接使用方便的 `Scan` 方法来直接读取一行，每次调用
    // 该方法可以让 scanner 读取下一行。
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        // `Text` 返回当前的 token，现在是输入的下一行。
        ucl := strings.ToUpper(scanner.Text())

        // 写出大写的行。
        fmt.Println(ucl)
    }

    // 检查 `Scan` 的错误。文件结束符是可以接受的，并且
    // 不会被 `Scan` 当作一个错误。
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
