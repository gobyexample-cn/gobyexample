// 想到等待多个协程完成，我们可以使用 *wait group*。

package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个协程都会运行该函数。
// 注意，WaitGroup 必须通过指针传递给函数。
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	// 睡眠一秒钟，以此来模拟耗时的任务。
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	// 通知 WaitGroup ，当前协程的工作已经完成。
	wg.Done()
}

func main() {

	// 这个 WaitGroup 被用于等待该函数开启的所有协程。
	var wg sync.WaitGroup

	// 开启几个协程，并为其递增 WaitGroup 的计数器。
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// 阻塞，直到 WaitGroup 计数器恢复为 0，即所有协程的工作都已经完成。
	wg.Wait()
}
