// Go 支持通过基于描述模板的时间格式化与解析。

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 这是一个遵循 RFC3339，
	// 并使用对应的 `布局`（layout）常量进行格式化的基本例子。
	t := time.Now()
	p(t.Format(time.RFC3339))

	// 时间解析使用与 `Format` 相同的布局值。
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` 和 `Parse` 使用基于例子的布局来决定日期格式，
	// 一般你只要使用 `time` 包中提供的布局常量就行了，但是你也可以实现自定义布局。
	// 布局时间必须使用 `Mon Jan 2 15:04:05 MST 2006` 的格式，
	// 来指定 格式化/解析给定时间/字符串 的布局。
	// 时间一定要遵循：2006 为年，15 为小时，Monday 代表星期几等规则。
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// 对于纯数字表示的时间（时间戳），
	// 您还可以将标准字符串格式与提取时间值的一部分一起使用。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// 当输入的时间格式不正确时，`Parse` 会返回一个解析错误。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
