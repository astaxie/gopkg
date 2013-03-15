## func LastIndexAny(s []byte, chars string) int

参数列表

- s 字节切片
- chars 字符串

返回值

- int s的位置索引（从0开始）

功能说明

- LastIndexAny把s解释为UTF-8编码的字节序列，返回chars中任意字符在s中最后出现的位置索引。如果chars为空或者s中不包含chars中的任意字符，则返回-1。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("123456789")
		fmt.Println(bytes.LastIndexAny(s, "789"))
		fmt.Println(bytes.LastIndexAny(s, "987"))
		fmt.Println(bytes.LastIndexAny(s, "0"))
		fmt.Println(bytes.LastIndexAny(s, ""))
	}

代码输出

	8
	8
	-1
	-1

	
