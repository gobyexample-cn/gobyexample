// Go 提供内建的 JSON 编码解码（序列化反序列化）支持，
// 包括内建及自定义类型与 JSON 数据之间的转化。

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 下面我们将使用这两个结构体来演示自定义类型的编码和解码。
type response1 struct {
	Page   int
	Fruits []string
}

// 只有 `可导出` 的字段才会被 JSON 编码/解码。必须以大写字母开头的字段才是可导出的。
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// 首先我们来看一下基本数据类型到 JSON 字符串的编码过程。
	// 这是一些原子值的例子。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 这是一些切片和 map 编码成 JSON 数组和对象的例子。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON 包可以自动的编码你的自定义类型。
	// 编码的输出只包含可导出的字段，并且默认使用字段名作为 JSON 数据的键名。
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 你可以给结构字段声明标签来自定义编码的 JSON 数据的键名。
	// 上面 `Response2` 的定义，就是这种标签的一个例子。
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 现在来看看将  JSON 数据解码为 Go 值的过程。
	// 这是一个普通数据结构的解码例子。
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 我们需要提供一个 JSON 包可以存放解码数据的变量。
	// 这里的 `map[string]interface{}` 是一个键为 string，值为任意值的 map。
	var dat map[string]interface{}

	// 这是实际的解码，以及相关错误的检查。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 为了使用解码 map 中的值，我们需要将他们进行适当的类型转换。
	// 例如，这里我们将 `num` 的值转换成 `float64` 类型。
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套的值需要一系列的转化。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们还可以将 JSON 解码为自定义数据类型。
	// 这样做的好处是，可以为我们的程序增加附加的类型安全性，
	// 并在访问解码后的数据时不需要类型断言。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 在上面例子的标准输出上，
	// 我们总是使用 byte和 string 作为数据和 JSON 表示形式之间的中介。
	// 当然，我们也可以像 `os.Stdout` 一样直接将 JSON 编码流传输到
	// `os.Writer` 甚至 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
