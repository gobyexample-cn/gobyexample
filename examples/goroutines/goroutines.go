// _协程(goroutine)_ 是轻量级的执行线程。

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 假设我们有一个函数叫做 `f(s)`。
	// 我们一般会这样 `同步地` 调用它
	f("direct")

	// 使用 `go f(s)` 在一个协程中调用这个函数。
	// 这个新的 Go 协程将会 `并发地` 执行这个函数。
	go f("goroutine")

	// 你也可以为匿名函数启动一个协程。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在两个协程在独立的协程中 `异步地` 运行，
	// 然后等待两个协程完成（更好的方法是使用 [WaitGroup](waitgroups)）。
	time.Sleep(time.Second)
	fmt.Println("done")
}
