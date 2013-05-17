# func (t Time) Local() Time

参数列表：

- 无

返回值：

- Time

功能说明：

返回本地时间。等价于t.In(time.Local)

代码实例：

	package main
	
	import (
	    "fmt"
	    "time"
	)
	
	func main() {
	    t := time.Date(2013, time.March, 14, 15, 23, 4, 100, time.UTC)
	    fmt.Println(t)
	    fmt.Println(t.Local())
	}
