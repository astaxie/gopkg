## func Count(s, sep []byte) int

参数列表

- s 字节切片
- sep 子字节切片

返回值

- int 子字节切片sep的数量

功能说明

- Count计算子字节切片sep在字节切片s中非重叠实例的数量

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("1234567777777")
		fmt.Println(bytes.Count(s, []byte("123")))
		fmt.Println(bytes.Count(s, []byte("77")))
		fmt.Println(bytes.Count(s, []byte("777")))
	}

	
代码输出

	1
	3
	2
