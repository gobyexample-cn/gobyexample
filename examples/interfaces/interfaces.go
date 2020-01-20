// 方法签名的集合叫做：_接口(Interfaces)_。

package main

import (
	"fmt"
	"math"
)

// 这是一个几何体的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

// 在这个例子中，我们将为 `rect` 和 `circle` 实现该接口。
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。
// 这里我们为 `rect` 实现了 `geometry` 接口。
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// `circle` 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量实现了某个接口，我们就可以调用指定接口中的方法。
// 这儿有一个通用的 `measure` 函数，我们可以通过它来使用所有的 `geometry`。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// 结构体类型 `circle` 和 `rect` 都实现了 `geometry` 接口，
	// 所以我们可以将其实例作为 `measure` 的参数。
	measure(r)
	measure(c)
}
