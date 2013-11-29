# func Tick(d Duration) <-chan Time

参数列表：

- d 时间跨度

返回值：

- 返回Time chan

功能说明：

Tick 是对NewTicker更易用的包装，提供了对定时器channel的访问。通常用于没必要停止定时器的客户端。

代码实例：

    package main

    import (
        "fmt"
    	"time"
    )
    
    func main() {
    	c := time.Tick(1 * time.Second)
    	for now := range c {
    		fmt.Printf("%v\n", now)
    	}
    }
