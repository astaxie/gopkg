## func (b *Writer) WriteByte(c byte) error

参数说明

- c 要写入的字节

返回值

- error 错误

功能说明

- 写入一个字节

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
		w.WriteByte('a')
		w.Flush()
		fmt.Println(string(wb.Bytes()))
	}

代码输出

	a
