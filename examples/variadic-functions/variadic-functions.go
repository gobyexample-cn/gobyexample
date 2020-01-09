// [_可变参数函数_](https://zh.wikipedia.org/wiki/%E5%8F%AF%E8%AE%8A%E5%8F%83%E6%95%B8%E5%87%BD%E6%95%B8)。
// 在调用时可以传递任意数量的参数。
// 例如，`fmt.Println` 就是一个常见的变参函数。

package main

import "fmt"

// 这个函数接受任意数量的 `int` 作为参数。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// 变参函数使用常规的调用方式，传入独立的参数。
	sum(1, 2)
	sum(1, 2, 3)

	// 如果你有一个含有多个值的 slice，想把它们作为参数使用，
	// 你需要这样调用 `func(slice...)`。
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
