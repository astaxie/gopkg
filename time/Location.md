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


# type Location struct


Location映射了时间瞬间和时间所在的区域.通常Location表示有相同时间偏移的地理区域的集合，例如中欧的CEST和CET。

	var Local *Location = &localLoc
	
Local表示系统的本地时区

	var UTC *Location = &utcLoc
	
UTC 表示通用协调时间。