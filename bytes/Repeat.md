## func Repeat(b []byte, count int) []byte

参数列表

- b 字节切片
- count 重复次数

返回值

- []byte 新建的字节切片

功能说明

- Repeat把b复制count次组合成一个新的字节切片并返回。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		b := []byte("123")
		fmt.Println(string(bytes.Repeat(b, 3)))
	}

代码输出

	123123123
