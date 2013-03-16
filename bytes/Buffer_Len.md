## func (b *Buffer) Len() int

返回值

- int 字节数

功能说明

- 返回Buffer中未读数据的字节数。它满足条件b.Len() == len(b.Bytes())。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("123456"))
		fmt.Println(b.Len())
		var buff [3]byte
		b.Read(buff[:])
		fmt.Println(b.Len())
	}
	
代码输出

	6
	3
