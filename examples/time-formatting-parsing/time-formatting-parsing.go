// Go 支持通过基于模板的描述进行时间格式化和解析。

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // 这里是一个基本的根据 RFC3339 进行格式化时间的例子。
    t := time.Now()
    p(t.Format("2006-01-02T15:04:05Z07:00"))

    // `Format` 使用一种基于例子的方式，使用一个已格式化的引
    // 用时间 `Mon Jan 2 15:04:05 MST 2006` 来决定用来为给定
    // 时间进行格式化的形式。例子中的时间和看到的的一样：2006
    // 表示年，15 是小时，Monday 是星期等等。这里是一些时间
    // 格式化的例子。
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))

    // 对于纯数字表示的时间，你也可以使用标准的格式化字
    // 符串来提出出时间值得组成。
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    // 时间解析使用的的相同的基于例子的方法。这个解析例
    // 子使用和上面一样的模式。
    withNanos := "2006-01-02T15:04:05.999999999-07:00"
    t1, e := time.Parse(
        withNanos,
        "2012-11-01T22:08:41.117442+00:00")
    p(t1)
    kitchen := "3:04PM"
    t2, e := time.Parse(kitchen, "8:41PM")
    p(t2)

    // `Parse` 函数在输入的时间格式不正确是会返回一个
    // 错误。
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)

    // 你可以在格式化和解析时间时使用一些预定义的时间格式。
    p(t.Format(time.Kitchen))
}
