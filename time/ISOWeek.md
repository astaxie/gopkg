# func (t Time) ISOWeek() (year, week int)

参数列表：

- 无

返回值：

- year 年
- week 星期

功能说明：

返回时间t的在一年中的星期数。星期范围从1到53。每年的1月1日到3日可能会属于上一年的第52或者53周，每年的12月29日到31日可能属于下一年的第一周。

代码实例：

	package main
	
	import (
	    "fmt"
	    "time"
	)
	
	func main() {
	    t := time.Date(2013, time.January, 7, 15, 23, 4, 100, time.UTC)
	    fmt.Println(t.ISOWeek())
	}