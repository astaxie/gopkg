# func type Time struct {}

精确到纳秒的时间

程序中通常使用值存储或传递，而不是指针。也就是说时间变量或结构体字段的类型应该是time.Time, 而不是*time.Time。一个Time值可以在多个不同的goroutine中同时使用。

Time实例可以通过Before，After和Equal方法进行比较。Sub方法可以使两个Time实例相减，得到Duration。Add方法使一个Time加上一个Duration得到新的Time。

Time类型的0值为January 1, year 1, 00:00:00.000000000 UTC。由于这个时间不太可能会在实践中用到，IsZero方法可以简单的判断一个Time是否被明确的初始化了。

每个Time和一个Location联系在一起，当计算时间的表现形式的时候起作用，例如Format，Hour，Year方法。Local，UTC和Ln方法，返回指定location的Time。这种形式改变location只是改变表现形式，并没有改变实际的时间，因此不影响之前的计算和描述。

示例代码：

	package main
	
	import (
	    "fmt"
	    "time"
	)
	
	func main() {
	    t1 := time.Date(2012, time.March, 12, 17, 45, 20, 30, time.Local)
	    fmt.Println(t1)
	    t2 := time.Now()
	    fmt.Println(t2)
	
	    fmt.Println(t1.After(t2), t1.Before(t2), t1.Equal(t2))
	    fmt.Println("Year =", t1.Year())
	    fmt.Println("Month =", t1.Month())
	    fmt.Println("Day =", t1.Day())
	    fmt.Println("Hour =", t1.Hour())
	    fmt.Println("Minute =", t1.Minute())
	    fmt.Println("Second =", t1.Second())
	    fmt.Println("Nanosecond =", t1.Nanosecond())
	    fmt.Println("Weekday =", t1.Weekday())
	
	    fmt.Println(t1.Clock())
	    fmt.Println(t1.Date())
	}