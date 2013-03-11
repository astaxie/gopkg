## func (b *Reader) Read(p []byte) (n int, err error)

参数列表：

- p 用于存放读取的数据的字节切片

返回值

- n 读取的字节数
- err 错误

功能说明

- 读取数据存放到p中，返回已读取的字节数。因为最多只调用底层的io.Reader一次，所以返回的n可能小于len(p)。在字节流结束时，n会为0并且err为io.EOF。

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
		var buf [128]byte
		n, err := r.Read(buf[:])
		fmt.Printf("%d, %v\n", n, err)
		fmt.Println(string(buf[:n]))
	}

代码输出
	
	8, <nil>
	12345678

	
