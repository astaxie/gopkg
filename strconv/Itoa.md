# func Itoa(i int) string

参数列表

- i     表示需要被转换为字符串的int值

返回值：

- 返回string 表示转换后的字符串。

功能说明：

- 将数值按照10进制转换为字符串,是等同于FormatInt(i, 10)的简写。

代码实例：

	package main
	
	import (
		"fmt"
		"strconv"
	)
	
    func main() {
        fmt.Println(strconv.FormatInt(9223372036854775807, 10))
    }

代码输出：

    9223372036854775807