// 在前面的例子中，我们了解了[生成外部进程](../spawning-processes/)
// 的知识，当我们需要访问外部进程时时需要这样做，但是有时候，我们只想
// 用其他的（也许是非 Go 程序）来完全替代当前的 Go 进程。这时候，我们
// 可以使用经典的 <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// 方法的 Go 实现。

package main

import "syscall"
import "os"
import "os/exec"

func main() {

    // 在我们的例子中，我们将执行 `ls` 命令。Go 需要提供我
    // 们需要执行的可执行文件的绝对路径，所以我们将使用
    // `exec.LookPath` 来得到它（大概是 `/bin/ls`）。
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    // `Exec` 需要的参数是切片的形式的（不是放在一起的一个大字
    // 符串）。我们给 `ls` 一些基本的参数。注意，第一个参数需要
    // 是程序名。
    args := []string{"ls", "-a", "-l", "-h"}

    // `Exec` 同样需要使用[环境变量](environment-variables.html)。
    // 这里我们仅提供当前的环境变量。
    env := os.Environ()

    // 这里是 `os.Exec` 调用。如果这个调用成功，那么我们的
    // 进程将在这里被替换成 `/bin/ls -a -l -h` 进程。如果存
    // 在错误，那么我们将会得到一个返回值。
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
