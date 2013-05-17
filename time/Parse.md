# func Parse(layout, value string) (Time, error)

参数列表：

- layout 格式
- value 表示时间的字符串

返回值：

- Time
- error

功能说明：

根据格式化的字符串返回一个它所代表的时间值，layout定义了标准时间的显示格式（Mon Jan 2 15:04:05 -0700 MST 2006），还用来描述了待分析的字符串。预定义的layout有ANSIC，UnixDate等（见[Contants](Contants.md)）。想多了解格式和标准时间定义的参见ANSIC文档。
value中省略的元素默认为0，或者不可能出现0时为1。比如转换“3：04pm”返回的时间为Jan 1, year 0, 15:04:00 UTC。年必须在0000..9999之间。星期会做语法检查，但是值会被忽略。

代码实例：

	package main
	
	import (
	    "fmt"
	    "time"
	)
	
	func main() {
	    t, _ := time.Parse(time.ANSIC, "Mon May 8 16:19:35 2013")
	    fmt.Println(t)
	}
