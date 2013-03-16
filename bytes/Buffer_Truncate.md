## func (b *Buffer) Truncate(n int)

参数列表

- n

功能说明

- Truncate只保留前n个字节的数据并清除其余数据。如果n < 0或者n > b.Len()则Truncate导致panic。n为0时，等同于b.Reset()。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("123456"))
		b.Truncate(3)
		fmt.Println(b.String())
	}
	
代码输出

	123
