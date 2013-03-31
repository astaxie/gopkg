## func Split(s, sep []byte) [][]byte

参数列表

- s 要分割的字节切片
- sep 用作分隔符的字节切片

返回值

- [][]byte 把s分割后的字节切片的切片

功能说明

- Split把s用sep分割成多个字节切片返回。如果sep为空，Split则把s切分成每个字节切片对应一个UTF-8字符。Split等效于参数n为-1的SplitN函数。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("你,好")
		for i, c := range bytes.Split(s, []byte{','}) {
			fmt.Printf("%d: %s(%d)\n", i, string(c), len(c))
		}
		for i, c := range bytes.Split(s, nil) {
			fmt.Printf("%d: %s(%d)\n", i, string(c), len(c))
		}
	}

代码输出

	0: 你(3)
	1: 好(3)
	0: 你(3)
	1: ,(1)
	2: 好(3)
