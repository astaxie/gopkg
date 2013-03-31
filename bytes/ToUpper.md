## func ToUpper(s []byte) []byte

参数列表

- s 字节切片

返回值

- []byte 被转换为字符大写的字节切片

功能说明

- ToUpper返回s的一个副本，并把其中所有的Unicode字符转换为大写。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("hello, world!")
		fmt.Println(string(bytes.ToUpper(s)))
		fmt.Println(string(s))
	}

代码输出

	HELLO, WORLD!
	hello, world!

