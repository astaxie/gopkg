## func LastIndex(s, sep []byte) int

参数列表

- s 字节切片
- sep  字节切片

返回值

- int s中的位置索引（从0开始）

功能说明

- LastIndex返回sep在s中最后一次出现的位置索引，如果s中不包含sep则返回-1。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("abcd,abcd")
		fmt.Println(bytes.LastIndex(s, []byte("abc")))
	}

代码输出

	5
