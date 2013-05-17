# func (t Time)UnixNano() int64

参数列表：

- 无

返回值：

- int64

功能说明：

返回t从January 1, 1970 UTC开始的Unix时间的纳秒数，如果纳秒数不能用int64来表示则结果是未定义的。注意：这意味着时间零点的UnixNano结果是未定义的。

代码实例：

	package main
	
	import (
	    "fmt"
	    "time"
	)
	
	func main() {
	    now := time.Now()
	    fmt.Println(now.UnixNano())
	}