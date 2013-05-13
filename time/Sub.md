# func (t Time) Sub(u Time) Duration

参数列表：

- u 

返回值：

- Duration

功能说明：

返回t - u

代码实例：

	package main
	
	import(
	    "fmt"
	    "time"
	)
	
	func main() {
	    t1 := time.Now()
	    time.Sleep(3 * time.Second)
	    t2 := time.Now()
	    fmt.Println(t2.Sub(t1))
	}