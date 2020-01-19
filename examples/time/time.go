// Go 为时间（time）和时间段（duration）提供了大量的支持；这儿有是一些例子。

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 从获取当前时间时间开始。
	now := time.Now()
	p(now)

	// 通过提供年月日等信息，你可以构建一个 `time`。
	// 时间总是与 `Location` 有关，也就是时区。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 你可以提取出时间的各个组成部分。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 支持通过 `Weekday` 输出星期一到星期日。
	p(then.Weekday())

	// 这些方法用来比较两个时间，分别测试一下是否为之前、之后或者是同一时刻，精确到秒。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 方法 `Sub` 返回一个 `Duration` 来表示两个时间点的间隔时间。
	diff := now.Sub(then)
	p(diff)

	// 我们可以用各种单位来表示时间段的长度。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 你可以使用 `Add` 将时间后移一个时间段，或者使用一个 `-` 来将时间前移一个时间段。
	p(then.Add(diff))
	p(then.Add(-diff))
}
