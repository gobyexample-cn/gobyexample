// Go 的_结构体(struct)_ 是带类型的字段(fields)集合。
// 这在组织数据时非常有用。

package main

import "fmt"

// 这里的 `person` 结构体包含了 `name` 和 `age` 两个字段。
type person struct {
	name string
	age  int
}

func main() {

	// 使用这个语法创建新的结构体元素。
	fmt.Println(person{"Bob", 20})

	// 你可以在初始化一个结构体元素时指定字段名字。
	fmt.Println(person{name: "Alice", age: 30})

	// 省略的字段将被初始化为零值。
	fmt.Println(person{name: "Fred"})

	// `&` 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40})

	// 使用`.`来访问结构体字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 也可以对结构体指针使用`.` - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)

	// 结构体是可变(mutable)的。
	sp.age = 51
	fmt.Println(sp.age)
}
