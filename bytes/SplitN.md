## func SplitN(s, sep []byte, n int) [][]byte

参数列表

- s 要分割的字节切片
- sep 用作分隔符的字节切片
- n 返回切片的长度

返回值

- [][]byte 把s分割后的字节切片的切片

功能说明

- SplitN把s用sep分割成多个字节切片返回。如果sep为空，Split则把s切分成每个字节切片对应一个UTF-8字符。
- 参数n决定返回切片的长度：
	+ n > 0: 最多返回n个子切片；最后可能子切片可能包含未切分的字节序列。
	+ n == 0: 返回空切片
	+ n < 0: 返回所有的子切片

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
		s := []byte("1,2,3,4")
		sep := []byte{','}
		print(bytes.SplitN(s, sep, 2))
		print(bytes.SplitN(s, sep, 0))
		print(bytes.SplitN(s, sep, -1))
	}

代码输出

	0: 1
	1: 2,3,4
	nil
	0: 1
	1: 2
	2: 3
	3: 4

