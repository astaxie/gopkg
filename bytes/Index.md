## func Index(s, sep []byte) int

参数列表

- s 要检查的字节切片
- sep 子字节切片

返回值

- int sep的位置索引（从0开始）

功能说明

- Index返回sep在s中第一次出现的位置索引，如果sep不在s中则返回-1

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("1234,12345678")
		fmt.Println(bytes.Index(s, []byte("34")))
		fmt.Println(bytes.Index(s, []byte("789")))
	}

代码输出

	2
	-1
