// 标准库的 `strings` 包提供了很多有用的字符串相关的函数。
// 这里是一些用来让你对这个包有个初步了解的例子。

package main

import s "strings"
import "fmt"

// 我们给 `fmt.Println` 一个短名字的别名，我们随后将会经常
// 用到。
var p = fmt.Println

func main() {

    // 这是一些 `strings` 中的函数例子。注意他们都是包中的
    // 函数，不是字符串对象自身的方法，这意味着我们需要考
    // 虑在调用时传递字符作为第一个参数进行传递。
    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    // 你可以在 [`strings`](http://golang.org/pkg/strings/)
    // 包文档中找到更多的函数

    // 虽然不是 `strings` 的一部分，但是仍然值得一提的是获
    // 取字符串长度和通过索引获取一个字符的机制。
    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}
