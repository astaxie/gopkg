# func (t Time) Unix() int64

参数列表：

- 无

返回值：

- int64

功能说明：

返回Unix时间——t相对于January 1, 1970 UTC的时间间隔

代码实例：

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		t := time.Now()
		fmt.Println(t.Unix())
	}
	
# func 