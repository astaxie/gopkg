## func (b *Writer) Available() int

返回值

- int 字节数

功能说明

- Available返回缓冲区中未使用的字节数

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
		fmt.Println(w.Available())
		w.WriteByte('1')
		fmt.Println(w.Available())
	}

代码输出

	4096
	4095
