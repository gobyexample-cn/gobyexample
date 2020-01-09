// Go 支持 <em><a href="http://zh.wikipedia.org/wiki/%E6%8C%87%E6%A8%99_(%E9%9B%BB%E8%85%A6%E7%A7%91%E5%AD%B8)">指针</a></em>，
// 允许在程序中通过 `引用传递` 来传递值和数据结构。

package main

import "fmt"

// 我们将通过两个函数：`zeroval` 和 `zeroptr` 来比较 `指针` 和 `值`。
// `zeroval` 有一个 `int` 型参数，所以使用值传递。
// `zeroval` 将从调用它的那个函数中得到一个实参的拷贝：ival。
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` 有一个和上面不同的参数：`*int`，这意味着它使用了 `int` 指针。
// 紧接着，函数体内的 `*iptr` 会 _解引用_ 这个指针，从它的内存地址得到这个地址当前对应的值。
// 对解引用的指针赋值，会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// 通过 `&i` 语法来取得 `i` 的内存地址，即指向 `i` 的指针。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针也是可以被打印的。
	fmt.Println("pointer:", &i)
}
