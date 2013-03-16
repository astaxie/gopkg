## func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)

参数列表

- w io.Writer

返回值

- n 写入字节数
- err 错误

功能说明

- WriteTo把Buffer中的数据写入w直到数据为空或者遇到错误，返回值n总是足够用int表示，使用int64类型是为了和io.WriterTo接口匹配。任何写入时遇到的错误都会被返回。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("12345"))
		w := bytes.NewBuffer(nil)
		n, err := b.WriteTo(w)
		fmt.Println(n, err)
		fmt.Println(string(w.Bytes()))
	}
	
代码输出
	
	5 <nil>
	12345
