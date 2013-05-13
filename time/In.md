# func (t Time) In(loc *Location) Time

参数列表：

- loc 位置

返回值：

- Time 

功能说明：

返回location设为loc的t。如果loc为nil则函数panic

代码实例：

	package main
	
	import (
	    "fmt"
	    "time"
	)
	
	func main() {
	    t := time.Now()
	    fmt.Println(t)
	    fmt.Println(t.In(time.UTC))
	}
