## func (b *Writer) Flush() error

返回值

- error 错误

功能说明

- Flush把缓冲区中的数据写入底层的io.Writer

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
		w.Write([]byte("123456"))
		fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
		w.Flush()
		fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
	}

代码输出

	0:
	6:123456
