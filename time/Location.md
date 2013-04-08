# func (t Time) Location() *Location

参数列表：

- 无

返回值：

- *Location

功能说明：

返回时间t的时区信息。

代码实例：

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		fmt.Println(time.Now().Location())
	}
