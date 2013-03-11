## func (b *Reader) Peek(n int) ([]byte, error)

参数列表

- n 要读取的字节数

返回值：

- 已读取的字节切片
- 错误

功能说明：

- 读取指定字节数的数据，这些被读取的数据不会从缓冲区中清除。在下次读取之后，本次返回的字节切片会失效。如果Peek返回的字节数不足n字节，则会同时返回一个错误说明原因。如果n比缓冲区要大，则错误为ErrBufferFull。

- 示例代码：

	package main

	import (
		"bytes"
		"bufio"
		"fmt"
	)

	func main() {
	}

代码输出：

	
