## func (b *Buffer) Bytes() []byte

返回值

- []byte 字节切片

功能说明

- Bytes返回Buffer中未读数据的字节切片，满足len(b.Bytes()) == b.Len()。如果调用者修改了返回字节切片的内容，Buffer中的未读数据也会被修改。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("0123456"))
		unread := b.Bytes()
		fmt.Println(string(unread))
		for i, c := range unread {
			unread[i] = 'A' + c - '0'
		}
		var buff [7]byte
		b.Read(buff[:])
		fmt.Println(string(buff[:]))
	}
	
代码输出

	0123456
	ABCDEFG
