// _range_ 用于迭代各种各样的数据结构。
// 让我们来看看如何在我们已经学过的数据结构上使用 `range`。

package main

import "fmt"

func main() {

	// 这里我们使用 `range` 来对 slice 中的元素求和。
	// 数组也可以用这种方法初始化并赋初值。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` 在数组和 slice 中提供对每项的索引和值的访问。
	// 上面我们不需要索引，所以我们使用 _空白标识符_ `_` 将其忽略。
	// 实际上，我们有时候是需要这个索引的。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` 在 map 中迭代键值对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` 也可以只遍历 map 的键。
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` 在字符串中迭代 unicode 码点(code point)。
	// 第一个返回值是字符的起始字节位置，然后第二个是字符本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
