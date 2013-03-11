# func Since(t Time) Duration

参数列表：

- t 时间

返回值：

- Duration

功能说明：

返回从时间t到当前时间的间隔，time.Now().Sub(t)的简写

代码实例：

    package main
    
    import (
      "fmt"
    	"time"
    )
    
    func main() {
    	t := time.Now()
    	time.Sleep(3 * time.Second)
    	fmt.Println(time.Since(t))
    }
