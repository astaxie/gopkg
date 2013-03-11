# func ParseDuration(s string) (Duration, error)

参数列表：
- s 时间长短的字符串

返回值：
- Duration 解析后的时间长短
- error 错误

功能说明：

将一个表示时间的字符串解析，每一个字符串是带有单位的十进制数序列，每个数字带有可选的单位或小数位。
合法的单位是"ns", "us" (or "µs"), "ms", "s", "m", "h"。

代码实例：

    package main

    import (
      "fmt"
    	"time"
    )
    
    func main() {
    	d, err := time.ParseDuration("1h2m3s4ms5.8us6ns")
    	fmt.Printf("%v %v", d, err)
    }
