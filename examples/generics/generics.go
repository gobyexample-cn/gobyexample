// 从1.18版本开始，Go添加了对泛型的支持，也即类型参数。

package main

import "fmt"

// 作为泛型函数的示例，`MapKeys` 接受任意类型的Map并返回其Key的切片。
// 这个函数有2个类型参数 - `K` 和 `V`；
// `K` 是 `comparable` 类型，也就是说我们可以通过 `==` 和 `!=`
// 操作符对这个类型的值进行比较。这是Go中Map的Key所必须的。
// `V` 是 `any` 类型，意味着它不受任何限制 (`any` 是 `interface{}` 的别名类型).
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// 作为泛型类型的示例， `List` 是一个
// 具有任意类型值的单链表。
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// 我们可以像在常规类型上一样定义泛型类型的方法
// 但我们必须保留类型参数。
// 这个类型是 `List[T]`，而不是 `List`
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// 当调用泛型函数的时候, 我们经常可以使用类型推断。
	// 注意，当调用 `MapKeys` 的时候，
	// 我们不需要为 `K` 和 `V` 指定类型 - 编译器会进行自动推断
	fmt.Println("keys m:", MapKeys(m))

	// ... 虽然我们也可以明确指定这些类型。
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
