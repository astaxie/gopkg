# func CanBackquote(s string) bool

参数列表

- s 字符串 

返回值：

- i 表示转换后的数值 

功能说明：

- 判断字符串 s 是否可以表示为一个单行的“反引号”字符串， 字符串中不能含有控制字符（除了 \t）和“反引号”字符，否则返回 false

代码实例：

	package main
	
	import (
		"fmt"
		"strconv"
	)
	
    func main() {
        s := strconv.CanBackquote("C:\\Windows\n")
        fmt.Println(s) // false
        s = strconv.CanBackquote("C:\\Windows\r")
        fmt.Println(s) // false
        s = strconv.CanBackquote("C:\\Windows\f")
        fmt.Println(s) // false
        s = strconv.CanBackquote("C:\\Windows\t")
        fmt.Println(s) // true
        s = strconv.CanBackquote("C:\\`Windows`")
        fmt.Println(s) // false
    }


代码输出：

    false
    false
    false
    true
    false