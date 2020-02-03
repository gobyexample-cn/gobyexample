// 在前面的例子中，我们用 [互斥锁](mutexes) 进行了明确的锁定，
// 来让共享的 state 跨多个 Go 协程同步访问。
// 另一个选择是，使用内建协程和通道的同步特性来达到同样的效果。
// Go 共享内存的思想是，通过通信使每个数据仅被单个协程所拥有，即通过通信实现共享内存。
// 基于通道的方法与该思想完全一致！

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 在这个例子中，state 将被一个单独的协程拥有。
// 这能保证数据在并行读取时不会混乱。
// 为了对 state 进行读取或者写入，
// 其它的协程将发送一条数据到目前拥有数据的协程中，
// 然后等待接收对应的回复。
// 结构体 `readOp` 和 `writeOp` 封装了这些请求，并提供了响应协程的方法。
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// 和前面的例子一样，我们会计算操作执行的次数。
	var readOps uint64
	var writeOps uint64

	// 其他协程将通过 `reads` 和 `writes` 通道来发布 `读` 和 `写` 请求。
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 这就是拥有 `state` 的那个协程，
	// 和前面例子中的 map 一样，不过这里的 state 是被这个状态协程私有的。
	// 这个协程不断地在 `reads` 和 `writes` 通道上进行选择，并在请求到达时做出响应。
	// 首先，执行请求的操作；然后，执行响应，在响应通道 `resp` 上发送一个值，表明请求成功（`reads` 的值则为 state 对应的值）。
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动 100 个协程通过 `reads` 通道向拥有 state 的协程发起读取请求。
	// 每个读取请求需要构造一个 `readOp`，发送它到 `reads` 通道中，
	// 并通过给定的 `resp` 通道接收结果。
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 用相同的方法启动 10 个写操作。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 让协程们跑 1s。
	time.Sleep(time.Second)

	// 最后，获取并报告 `ops` 值。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
