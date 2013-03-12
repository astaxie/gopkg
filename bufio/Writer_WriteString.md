## func (b *Writer) WriteString(s string) (int, error)

参数说明

- s 要写入的字符串

返回值

- int 写入字节数
- error 错误

功能说明

- WriteString写入一个字符串，返回写入的字节数。如果返回的字节数小于len(s)，则同时返回一个错误说明原因。

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
		n, err := w.WriteString("123456")
		if err != nil {
			return
		}
		fmt.Println(n)
	}

代码输出

	6

	
