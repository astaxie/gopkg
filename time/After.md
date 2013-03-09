## func After(d Duration) <-chan Time
参数列表：
- d 时间跨度

返回值：
- 返回Time类型chan

功能说明：

这个函数用来实现指定时间的延时，返回一个单向的读取chan

代码实例：
    
    package main

    import (
        "fmt"
        "time"
    )
    
    func main() {
    	result := make(chan int)
    	go func(ch chan int) {
    		time.Sleep(3 * time.Second)	
    		ch <- 4
    	}(result)
    
    	select {
    	case <-time.After(2 * time.Second):
    		fmt.Println("Time out")
    	case <-result:
    		fmt.Println(result)
    	}
    }
