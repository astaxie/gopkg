## func Compare(a, b []byte) int

参数列表

- a 第一个要比较的字节切片
- b 第二个要比较的字节切片

返回值

- int 如果a==b返回0，如果a < b返回-1，如果a > b返回1

功能说明

- Compare根据字节的值比较两个字节切片的大小

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		fmt.Println(bytes.Compare([]byte{1}, []byte{2}))
		fmt.Println(bytes.Compare([]byte{1}, []byte{1}))
		fmt.Println(bytes.Compare([]byte{1}, []byte{0}))
		fmt.Println(bytes.Compare([]byte{1}, []byte{1, 2}))
		fmt.Println(bytes.Compare([]byte{2}, []byte{1, 2}))
	}

代码输出
	
	-1
	0
	1
	-1
	1

