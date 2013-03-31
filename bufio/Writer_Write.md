## func (b *Writer) Write(p []byte) (nn int, err error)

参数说明

- p 要写入的字节切片

返回值

- nn 已写的字节数
- err 错误

功能说明

- Write把p写入缓冲区，返回已写入的字节数。如果nn小于len(p)，则同时返回一个错误说明原因。

代码示例

	package main

	import (
		"bufio"
		"bytes"
		"fmt"
	)

	func main() {
		wb := bytes.NewBuffer(nil)
		w := bufio.NewWriter(wb)
		n, err := w.Write([]byte("123456"))
		if err != nil {
			return
		}
		fmt.Println(n)
		w.Flush()
		fmt.Println(string(wb.Bytes()))
	}

代码输出

	6
	123456
