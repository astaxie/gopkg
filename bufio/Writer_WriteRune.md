## func (b *Writer) WriteRune(r rune) (size int, err error)

参数说明

- r 要写入的Unicode字符

返回值

- size 写入的字节数
- err 错误

功能说明

- WriteRune以UTF-8编码写入一个Unicode字符，返回写入的字节数和错误。

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
		n, err := w.WriteRune('你')
		if err != nil {
			return
		}
		w.Flush()
		fmt.Println(n)
		fmt.Println(wb.Bytes())
	}

代码输出

	3
	[228 189 160]
