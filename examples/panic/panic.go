// `panic` 意味着有些出乎意料的错误发生。通常我们用它
// 来表示程序正常运行中不应该出现的，后者我么没有处理
// 好的错误。

package main

import "os"

func main() {

    // 我们将在真个网站中使用 panic 来检查预期外的错误。这个
    // 是唯一一个为 panic 准备的例子。
    panic("a problem")

    // panic 的一个基本用法就是在一个函数返回了错误值但是我们并不知道（或
    // 者不想）处理时终止运行。这里是一个在创建一个新文件时返回异常错误时的
    // `panic` 用法。
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
