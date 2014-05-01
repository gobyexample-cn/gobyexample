// Go 支持通过基于描述模板的时间格式化和解析。

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // 这里是一个基本的按照 RFC3339 进行格式化的例子，使用
    // 对应模式常量。
    t := time.Now()
    p(t.Format(time.RFC3339))

    // 时间解析使用同 `Format` 相同的形式值。
    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)

    // `Format` 和 `Parse` 使用基于例子的形式来决定日期格式，
    // 一般你只要使用 `time` 包中提供的模式常量就行了，但是你
    // 也可以实现自定义模式。模式必须使用时间 `Mon Jan 2 15:04:05 MST 2006`
    // 来指定给定时间/字符串的格式化/解析方式。时间一定要按照
    // 如下所示：2006为年，15 为小时，Monday 代表星期几，等等。
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)

    // 对于纯数字表示的时间，你也可以使用标准的格式化字
    // 符串来提出出时间值得组成。
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    // `Parse` 函数在输入的时间格式不正确是会返回一个
    // 错误。
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}
