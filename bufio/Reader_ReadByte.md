## func (b *Reader) ReadByte() (c byte, err error)

返回值

- c 读取的字节
- err 错误

功能说明

- 读取并返回一个字节。如果没有字节可读，则返回错误。

代码示例

	package main

	import (
		"bytes"
		"bufio"
		"fmt"
	)

	func main() {
		rb := bytes.NewBuffer([]byte("12345678"))
		r := bufio.NewReader(rb)
		b, err := r.ReadByte()
		fmt.Printf("%c, %v\n", b, err)
	}

代码输出
	
	1, <nil>

	
