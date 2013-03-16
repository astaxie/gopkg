## func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)

参数列表

- delim 分隔字节

返回值

- line 字节切片
- err 错误

功能说明

- ReadBytes读取数据直到第一次遇到delim，把已读取的数据包括delim作为字节切片返回。如果在读取到delim前出现错误，则返回已读取的数据和那个错误（通常是io.EOF）。只有返回的数据不以delim结尾时，ReadBytes才返回err非空。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("123,456"))
		line, err := b.ReadBytes(',')
		fmt.Println(string(line))
		fmt.Println(err)
	}
	
代码输出

	123,
	<nil>
