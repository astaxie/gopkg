# func AfterFunc(d Duration, f func()) *Timer

参数列表：

- d 时间间隔
- f 回调函数

返回值：

- *Timer 

功能说明：

经过时间d之后在自己的goroutine中调用f。可以调用返回的Timer的Stop方法来停止Timer。

代码实例：

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		f := func() {
			fmt.Println("Time out")
		}
		time.AfterFunc(2*time.Second, f)
		
		time.Sleep(3 * time.Second)
	}
