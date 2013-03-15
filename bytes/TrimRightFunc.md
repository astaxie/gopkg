## func TrimRightFunc(s []byte, f func(r rune) bool) []byte

参数列表

- s 字节切片
- f 检查函数

返回值

- []byte s的子字节切片

功能说明

- TrimRightFunc返回s的一个子字节切片，不包含s尾部连续满足f(c)==true的字符c。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("123456789")
		f := func(r rune) bool {
			return r > '7'
		}
		fmt.Println(string(bytes.TrimRightFunc(s, f)))
	}

代码输出

	1234567
