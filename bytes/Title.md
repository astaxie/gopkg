## func Title(s []byte) []byte

参数列表

- s 字节切片

返回值

- []byte 被转换为首字母大写的字节切片

功能说明

- Title返回s的一个副本，把s中每个单词的首字母改为Unicode字符的大写。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("hello, world!")
		fmt.Println(string(bytes.Title(s)))
		fmt.Println(string(s))
	}

代码输出

	Hello, World!
	hello, world!
