# func FormatBool(b bool) string

参数列表

- b 需要被转换的bool值 

返回值：

- 返回string 表示转换后的字符串 

功能说明：

- 将true或false转换为字符串

代码实例：

	package main
	
	import (
		"fmt"
		"strconv"
	)
	
    func main() {
        b := strconv.FormatBool(true)
        fmt.Println(b)
        b = strconv.FormatBool(false)
        fmt.Println(b)
    }


代码输出：

    true
    false