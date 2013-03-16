## func (b *Buffer) WriteString(s string) (n int, err error)

参数列表

- s 要写入的字符串

返回值

- n 写入的字节数
- err 错误

功能说明

- WriteString把s写入Buffer，返回值n为len(s)，err总是为nil。如果内部缓冲区太大，WriteString会panic并产生ErrTooLarge异常。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer(nil)
		b.WriteString("你好世界")
		fmt.Println(string(b.Bytes()))
	}
	
代码输出

	你好世界
