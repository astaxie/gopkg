# func (t Time) Format(layout string) string

参数列表：

- layout 时间格式

返回值：

- string 时间格式化之后的字符串

功能说明：

返回根据layout指定的格式格式化之后的字符串，layout定义了标准时间的显示格式。预定义的layout有ANSIC，UnixDate，RFC3339等。

常量：

    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
    

代码实例：

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		t := time.Now()
		fmt.Println(t.Format(time.ANSIC))
		fmt.Println(t.Format(time.UnixDate))
		//...
	}
