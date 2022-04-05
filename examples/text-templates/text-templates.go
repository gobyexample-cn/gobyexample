// Go 使用 `text/template` 包为创建动态内容或向用户显示自定义输出提供了内置支持。
// 一个名为 `html/template` 的兄弟软件包提供了相同的 API，但具有额外的安全功能，被用于生成 HTML。

package main

import (
	"os"
	"text/template"
)

func main() {

	// 我们可以创建一个新模板，并从字符串解析其正文。
	// 模板是静态文本和包含在”动作“中用于动态插入内容的混合体。
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// 或者，我们可以使用 `template.Must` 函数，在 `Parse` 错误时返回 panic。
	// 这对于在全局作用域中初始化的模板非常有用。
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// 通过“执行”模板，我们生成其文本，其中包含其“动作”的具体值。
	// `{{.}}` 动作被作为参数传递给 `Execute` 的值所代替。
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// 我们将在下面使用辅助函数。
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// 如果数据是一个结构体，我们可以使用 `{{.FieldName}}` 动作来访问其字段。
	// 这些字段应该是导出的，以便在模板执行时可访问。
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// 这同样适用于 map；在 map 中没有限制键名的大小写。
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else 提供了条件执行模板。如果一个值是类型的默认值，例如 0、空字符串、空指针等，
	// 则该值被认为是 false。
	// 这个示例演示了另一个模板特性：使用 `-` 在动作中去除空格。
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// range 块允许我们遍历切片，数组，映射或通道。在 range 块内，`{{.}}` 动作被设置为迭代的当前项。
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
