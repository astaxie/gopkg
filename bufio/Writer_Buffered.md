## func (b *Writer) Buffered() int

返回值

- int 字节数

功能说明

- 返回已写入当前缓冲区中的字节数

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
		fmt.Println(w.Buffered())
		w.WriteByte('1')
		fmt.Println(w.Buffered())
		w.Flush()
		fmt.Println(w.Buffered())
	}

代码输出

	0
	1
	0


	
