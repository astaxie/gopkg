# func (d Duration) String() string

参数列表：

- 无

返回值：

- string

功能说明：

根据时间返回一个"72h3m0.5s"形式的字符串，前导0会被省略。
特殊情况，当时间小于1秒时，会用更小的单位（毫秒，微秒，纳秒）来保证前导数字不为0。时间就为0时则仅返回0，没有单位

代码实例：

    package main
    
    import (
      "fmt"
    	"time"
    )
    
    func main() {
    	d, _ := time.ParseDuration("1h2m3.5s")
    	fmt.Println(d)
    	d, _ = time.ParseDuration("3ms4ns")
    	fmt.Println(d)
    	d, _ = time.ParseDuration("0")
    	fmt.Println(d)
    }

#func (m Month) String() string

返回月份的英文形式

示例代码：

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		fmt.Println(time.Month(1))
	}
