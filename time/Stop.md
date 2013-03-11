# func (t *Ticker) Stop()

参数列表：

- 无

返回值：

- 无

功能说明：

停止一个定时器

代码实例：

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func tick(ch <-chan time.Time) {
		for t := range ch {
			fmt.Println(t)
		}
	}
	
	func main() {
		ticker := time.NewTicker(time.Second)
	
		go tick(ticker.C)
	
		<-time.After(5 * time.Second)
		ticker.Stop()
	}
