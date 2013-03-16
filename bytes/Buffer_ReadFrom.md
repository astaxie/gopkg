## func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)

参数列表

- r io.Reader

返回值

- n 已读取字节数
- err 错误

功能说明

- ReadFrom从r中读取数据写入Buffer直到r.Read返回io.EOF，除了io.EOF之外的错误会被ReadFrom返回。返回值n为已读取的字节数，err为r.Read返回的错误。如果读取的数据太大，ReadFrom会panic并产生ErrTooLarge异常。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		r := bytes.NewBuffer([]byte("123"))
		b := bytes.NewBuffer(nil)
		n, err := b.ReadFrom(r)
		fmt.Println(n, err)
		fmt.Println(string(b.Bytes()))
	}
	
代码输出

	3 <nil>
	123
