## func SplitAfter(s, sep []byte) [][]byte

参数列表

- s 要分割的字节切片
- sep 用作切分标志的字节切片

返回值

- [][]byte 切分后的字节切片的切片

功能说明

- SplitAfter用sep作为后缀把s切分成多个字节切片返回。如果sep为空，则把s切分成每个字节切片对应一个UTF-8字符。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func print(s [][]byte) {
		if len(s) == 0 {
			fmt.Println("nil")
			return
		}
		for i, c := range s {
			fmt.Printf("%d: %s\n", i, string(c))
		}
	}

	func main() {
		s := []byte("1,2,3")
		sep := []byte{','}
		print(bytes.SplitAfter(s, sep))
		print(bytes.SplitAfter(s, nil))
	}

代码输出

	0: 1,
	1: 2,
	2: 3
	0: 1
	1: ,
	2：2
	3：,
	4：3
