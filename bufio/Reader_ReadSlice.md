## func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

参数说明

- delim 分隔字节

返回值

- line 字节切片
- err 错误

功能说明

- ReadSlice读取数据直到delim出现，并返回读取数据的字节切片。下次读取数据时返回的切片会失效。如果ReadSlice在查找到delim之前遇到错误，它返回读取的所有数据和那个错误（通常是io.EOF）。如果缓冲区满时也没有查找到delim，则返回ErrBufferFull错误。因为ReadSlice返回的数据会在下次I/O操作时被覆盖，大多数调用者应该使用ReadBytes或者ReadString。只有当line不以delim结尾时，ReadSlice才会返回非空err。

代码示例

	package main

	import (
		"bufio"
		"bytes"
		"fmt"
	)

	func main() {
		rb := bytes.NewBuffer([]byte("1234$56789"))
		r := bufio.NewReader(rb)
		line, err := r.ReadSlice('$')
		if err != nil {
			return
		}
		fmt.Println(string(line))
		r.ReadSlice('$')
		fmt.Println(string(line))
	}

代码输出

	1234$
	56789
