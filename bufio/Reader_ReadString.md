## func (b *Reader) ReadString(delim byte) (line string, err error)

参数说明

- delim 分隔字节

返回值

- line 返回的字符串
- err 错误

功能说明

- ReadString读取数据直到delim第一次出现，返回一个包含delim的字符串。如果ReadString在读取到delim前遇到错误，它返回已读字符串和那个错误（通常是io.EOF）。只有当返回的字符串不以delim结尾时，ReadString才返回非空err。

代码示例

	package main

	import (
		"bufio"
		"bytes"
		"fmt"
	)

	func main() {
		rb := bytes.NewBuffer([]byte("1234$5678"))
		r := bufio.NewReader(rb)
		line, err := r.ReadString('$')
		if err == nil {
			fmt.Println(line)
		}
	}

代码输出
	
	1234$
