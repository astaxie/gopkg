## func ToLower(s []byte) []byte

参数列表

- s 字节切片

返回值

- []byte 被转换为字符小写的字节切片

功能说明

- ToLower返回s的一个副本，并把其中所有的Unicode字符转换为小写。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("Hello, World!")
		fmt.Println(string(bytes.ToLower(s)))
		fmt.Println(string(s))
	}

代码输出

	hello, world!
	Hello, World!

