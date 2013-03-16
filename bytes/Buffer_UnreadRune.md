## func (b *Buffer) UnreadRune() error

返回值

- err 错误

功能说明

- UnreadRune取消上次ReadRune读取的Unicode字符。如果最近一次读写操作不是ReadRune，则UnreadRune返回一个错误。

代码示例

  	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewBuffer([]byte("你好世界"))
		r, _, _ := b.ReadRune()
		fmt.Println(string(r))
		fmt.Println(string(b.Bytes()))
		b.UnreadRune()
		fmt.Println(string(b.Bytes()))
	}
	
代码输出

	你
	好世界
	你好世界
