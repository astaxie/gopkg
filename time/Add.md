# func (t Time) Add(d Duration) Time

参数列表：

- d 时间间隔

返回值：

- Time

功能说明：

返回时间（t + d）

代码实例：

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		now := time.Now()
		fmt.Println("now:", now)
		fmt.Println("after 3 hours", now.Add(3*time.Hour))
	}
