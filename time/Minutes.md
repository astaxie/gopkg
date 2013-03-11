# func (d Duration) Minutes() float64

参数列表：

- 无

返回值：

- float64

功能说明：

返回时间d以分钟为单位的浮点数形式

代码实例：

    package main
    
    import (
      "fmt"
      "time"
    )
    
    func main() {
    	d, _ := time.ParseDuration("3h4m2s")
    	fmt.Println(d.Minutes())
    }
