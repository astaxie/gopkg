## func (r *Reader) Len() int

返回值

- int Reader中未读数据的字节数

功能说明

- Len返回Reader中未读数据的字节数

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewReader([]byte("12345"))
		fmt.Println(b.Len())
	}

代码输出
	
	5
