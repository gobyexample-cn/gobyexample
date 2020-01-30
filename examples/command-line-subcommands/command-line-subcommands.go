// `go` 和 `git` 这种命令行工具，都有很多的 *子命令* 。
// 并且每个工具都有一套自己的 flag，比如：
// `go build` 和 `go get` 是 `go` 里面的两个不同的子命令。
// `flag` 包让我们可以轻松的为工具定义简单的子命令。

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// 我们使用 `NewFlagSet` 函数声明一个子命令，
	// 然后为这个子命令定义一个专用的 flag。
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// 对于不同的子命令，我们可以定义不同的 flag。
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// 子命令应作为程序的第一个参数传入。
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// 检查哪一个子命令被调用了。
	switch os.Args[1] {

	// 每个子命令，都会解析自己的 flag 并允许它访问后续的位置参数。
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
