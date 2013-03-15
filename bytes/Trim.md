## func Trim(s []byte, cutset string) []byte

参数列表

- s 字节切片
- cutset 字符串

返回值

- []byte s的子字节切片

功能说明

- Trim返回s的子字节切片，不包含s首部和尾部的连续的cutset中的任意字符。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("123456789")
		fmt.Println(string(bytes.Trim(s, "1389")))
	}

代码输出

	234567
