## func (b *Reader) Peek(n int) ([]byte, error)

参数列表

- n 字节数

返回值：

- []byte 读取到的字节切片
- error 错误

功能说明：

- 读取指定字节数的数据，这些被读取的数据不会从缓冲区中清除。在下次读取之后，本次返回的字节切片会失效。如果Peek返回的字节数不足n字节，则会同时返回一个错误说明原因。如果n比缓冲区要大，则错误为ErrBufferFull。

示例代码：

	package main

	import (
		"bytes"
		"bufio"
		"fmt"
	)

	func main() {
		rb := bytes.NewBuffer([]byte("12345678"))
		r := bufio.NewReader(rb)
		b1, _ := r.Peek(4)
		fmt.Println(string(b1))
		b2, _ := r.Peek(8)
		fmt.Println(string(b2))
	}

代码输出：

	1234
	12345678

