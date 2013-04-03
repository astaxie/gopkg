## func Replace(s, old, new []byte, n int) []byte

参数列表

- s 字节切片
- old 字节切片
- new 字节切片
- n 替换次数

返回值

- []byte 字节切片s的一个副本

功能说明

- Replace返回字节切片s的一个副本，并把前n个不重叠的old替换为new；如果n < 0则不限制替换的数量。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("1234,234,234")
		d := bytes.Replace(s, []byte("234"), []byte("5678"), 2)
		fmt.Println(string(d))
	}

代码输出

	15678,5678,234
