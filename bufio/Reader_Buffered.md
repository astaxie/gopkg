## func (b *Reader) Buffered() int

返回值：

- int 字节数

功能说明：

- 返回可从缓冲区读出数据的字节数

代码示例：

	package main

	import (
		"bytes"
		"bufio"
		"fmt"
	)

	func main() {
		rb := bytes.NewBuffer([]byte("12345678"))
		r := bufio.NewReader(rb)
		fmt.Println(r.Buffered())
		var buf [4]byte
		r.Read(buf[:])
		fmt.Println(r.Buffered())
		r.Read(buf[:])
		fmt.Println(r.Buffered())
	}

代码输出：

	0
	4
	0

