## func TrimSpace(s []byte) []byte

参数列表

- s 字节切片

返回值

- []byte

功能说明

- TrimSpace返回s的一个子字节切片，不包含s中开始和结尾处的连续的Unicode空白字符。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s1 := bytes.TrimSpace([]byte("   \t  hello, world!\r\n"))
		s2 := bytes.TrimSpace([]byte("   \t  \r\n"))
		fmt.Printf("%d %s\n", len(s1), string(s1))
		fmt.Printf("%d %s\n", len(s2), string(s2))
	}

代码输出

	13 hello, world!
	0
