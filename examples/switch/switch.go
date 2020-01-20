// _switch_ 是多分支情况时快捷的条件语句。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 一个基本的 `switch`。
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在同一个 `case` 语句中，你可以使用逗号来分隔多个表达式。
	// 在这个例子中，我们还使用了可选的 `default` 分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// 不带表达式的 `switch` 是实现 if/else 逻辑的另一种方式。
	// 这里还展示了 `case` 表达式也可以不使用常量。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 类型开关 (`type switch`) 比较类型而非值。可以用来发现一个接口值的类型。
	// 在这个例子中，变量 `t` 在每个分支中会有相应的类型。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
