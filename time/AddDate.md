# func (t Time) AddDate(years int, months int, days int) Time

参数列表：

- years 年
- months 月
- days 日

返回值：

- Time

功能说明：

生成由t加上参数中指定的年月日之后的时间，比如AddDate(-1, 2, 3) 使 January 1, 2011变为March 4, 2010。AddDate以和Date方法一样的方式标准化返回值，比如，10月31日加一个月之后为11月31日，标准化为12月1日

代码实例：

	package main
	
	import (
	    "fmt"
	    "time"
	)
	
	func main() {
	    now := time.Now()
	    fmt.Println(now.AddDate(2, 3, 4))
	}