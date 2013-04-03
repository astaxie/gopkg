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

# func (t *Timer) Stop() (ok bool)

参数列表：

- 无

返回值：

- ok 成功true，失败false

功能说明：

停止Timer触发，成功返回true，已经停止或过期返回false。

代码实例：

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func callback() {
		fmt.Println("Time out")
	}
	
	func main() {
		time.AfterFunc(2*time.Second, callback).Stop()
		time.Sleep(3 * time.Second)
		fmt.Println("end")
	}
