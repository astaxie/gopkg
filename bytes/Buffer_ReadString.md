## func (b *Buffer) ReadString(delim byte) (line string, err error)

参数列表

- delim 分隔字节

返回值

- line 字符串
- err 错误

功能说明

- ReadString读取数据直到遇到delim，并把已读取的数据包含delim作为string返回。如果在遇到delim前出错，则把已读数据作为string和遇到的错误（通常是io.EOF）返回。当返回的string不已delim结尾时，才会返回非空err。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("123*456"))
		line, err := b.ReadString('*')
		fmt.Println(line, err)
	}
	
代码输出

	123* <nil>
