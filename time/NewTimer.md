# func NewTimer(d Duration) *Timer

参数列表：

- d 时间间隔

返回值：

- *Timer

功能说明：

新建一个Timer，在时间d之后将当前时间发送给C

代码实例：

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		timer := time.NewTimer(2 * time.Second)
		
		t := <-timer.C
		fmt.Println(t)
	}