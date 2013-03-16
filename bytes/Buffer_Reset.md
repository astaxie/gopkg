## func (b *Buffer) Reset()

功能说明

- Reset把Buffer中的数据清空，等同于b.Truncate(0)。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("123"))
		fmt.Println(b.Len())
		b.Reset()
		fmt.Println(b.Len())
	}
	
代码输出

	3
	0
