## func TrimFunc(s []byte, f func(r rune) bool) []byte

参数列表

- s 字节切片
- f 检查函数

返回值

- []byte s的子字节切片

功能说明

- TrimFunc返回s的子字节切片，不包含s首部和尾部连接的满足f(c)==true的字符c。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("12345678")
		f := func(r rune) bool {
			return r <= '3' || r >= '6'	
		}
		fmt.Println(string(bytes.TrimFunc(s, f)))
	}

代码输出

	45
