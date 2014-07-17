# func Atoi(s string) (i int, err error)

参数列表

- s 表示数字字符串，如“1234” 

返回值：

- i 表示转换后的数值 

功能说明：

- Atoi是函数ParseInt(s, 10, 0)的简写。把字符串格式的数字如“12345”转化为数字12345

代码实例：

	package main
	
	import (
		"fmt"
		"strconv"
	)
	
    func main() {
        fmt.Println(strconv.Atoi("12345"))
        fmt.Println(strconv.Atoi("abcde"))
    }


代码输出：

    12345 <nil>
    0 strconv.ParseInt: parsing "abcde": invalid syntax