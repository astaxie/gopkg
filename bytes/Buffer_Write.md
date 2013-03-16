## func (b *Buffer) Write(p []byte) (n int, err error)

参数列表

- p 要写入的字节切片

返回值

- n 已写入的字节数
- err 错误

功能说明

- Write把p写入Buffer，返回的n满足n == len(p)，err总是为nil。如果数据太大则Write会panic并产生ErrTooLarge异常。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer(nil)
		fmt.Println(b.Write([]byte("12345")))
		fmt.Println(string(b.Bytes()))
	}
	
代码输出

	5 <nil>
	12345
