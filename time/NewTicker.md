# func NewTicker(d Duration) *Ticker

参数列表：

- d 时钟间隔

返回值：

- *Ticker 定时器

功能说明：

新建一个Ticker，包含了time channel，每隔指定d间隔的时间发送时间给这个channel。d必须大于0，否则函数会崩溃

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
	    time.Sleep(5 * time.Second)
	}