// Go 的 `encoding.xml` 包为 XML 和 类-XML 格式提供了内建支持。

package main

import (
    "encoding/xml"
    "fmt"
)

// 该类型将被映射为 XML。
// 与 JSON 例子类型，字段 tag 包含了编码和解码的指令。
// 这里我们使用了 XML 包的一些特性：
// `XMLName` 字段名规定了 struct 的 XML 元素的名称；
// `id,attr` 意思是 `Id` 字段是一个 XML _attribute_，而不是嵌套元素。
type Plant struct {
    XMLName xml.Name `xml:"plant"`
    Id      int      `xml:"id,attr"`
    Name    string   `xml:"name"`
    Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
        p.Id, p.Name, p.Origin)
}

func main() {
    coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}

    // 传入我们声明了 XML 的 plant 类型。
    // 使用 `MarshalIndent` 生成可读性更好的输出结果。
    out, _ := xml.MarshalIndent(coffee, " ", "  ")
    fmt.Println(string(out))

    // 明确的为输出结果添加一个通用的 XML 头部信息
    fmt.Println(xml.Header + string(out))

    // 使用 `Unmarshal` 将 XML 格式的字节流解析到 struct 内。
    // 如果 XML 格式不正确，或无法映射到 struct，将会返回一个描述性错误。
    var p Plant
    if err := xml.Unmarshal(out, &p); err != nil {
        panic(err)
    }
    fmt.Println(p)

    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}

    // `parent>child>plant` 字段标签告诉编码器嵌套 `<parent><child>...` 下面的所有 plant。
    type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"`
    }

    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}

    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))
}
