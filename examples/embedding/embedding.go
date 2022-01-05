// Go支持对于结构体(struct)和接口(interfaces)的 _嵌入(embedding)_
// 以表达一种更加无缝的 _组合(composition)_ 类型

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// 一个 `container` _嵌入_ 了一个 `base`.
// 一个嵌入看起来像一个没有名字的字段
type container struct {
	base
	str string
}

func main() {

	// 当创建含有嵌入的结构体，必须对嵌入进行显式的初始化；
	// 在这里使用嵌入的类型当作字段的名字
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// 我们可以直接在 `co` 上访问 base 中定义的字段,
	// 例如： `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// 或者，我们可以使用嵌入的类型名拼写出完整的路径。
	fmt.Println("also num:", co.base.num)

	// 由于 `container` 嵌入了 `base`，因此`base`的方法
	// 也成为了 `container` 的方法。在这里我们直接在 `co` 上
	// 调用了一个从 `base` 嵌入的方法。
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// 可以使用带有方法的嵌入结构来赋予接口实现到其他结构上。
	// 因为嵌入了 `base` ，在这里我们看到 `container`
	// 也实现了 `describer` 接口。
	var d describer = co
	fmt.Println("describer:", d.describe())
}
