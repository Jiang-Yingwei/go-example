package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(then)
	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())
	//周几
	fmt.Println(then.Weekday())
	//这些方法来比较两个时间，分别测试一下是否是之前，之后或者是同一时刻，精确到秒。
	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))
	fmt.Println(then.Equal(now))
	//方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间。
	diff := now.Sub(then)
	fmt.Println(diff)
	//我们计算出不同单位下的时间长度值。
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())
	//可以使用 Add 将时间后移一个时间间隔，或者使用一个 - 来将时间前移一个时间间隔。
	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))
	//分别使用带 Unix 或者 UnixNano 的 time.Now来获取从自协调世界时起到现在的秒数或者纳秒数。
	secs := now.Unix()
	nanos := now.UnixNano()
	//注意 UnixMillis 是不存在的，所以要得到毫秒数的话，你要自己手动的从纳秒转化一下。
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	//你也可以将协调世界时起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
	//基本的按照 RFC3339 进行格式化
	fmt.Println(now.Format(time.RFC3339))
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	fmt.Println(t1)
	//Format 和 Parse 使用基于例子的形式来决定日期格式，一般你只要使用 time 包中提供的模式常量就行了，但是你也可以实现自定义模式。
	//模式必须使用时间 Mon Jan 2 15:04:05 MST 2006来指定给定时间/字符串的格式化/解析方式。
	//时间一定要按照如下所示：2006为年，15 为小时，Monday 代表星期几，等等。
	fmt.Println(now.Format("3:04PM"))
	fmt.Println(now.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(now.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	fmt.Println(t2)
	//对于纯数字表示的时间，你也可以使用标准的格式化字符串来提出出时间值得组成。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	//Parse 函数在输入的时间格式不正确是会返回一个错误。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	fmt.Println(e)
}
