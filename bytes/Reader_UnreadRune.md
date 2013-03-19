## func (r *Reader) UnreadRune() error

返回值

- error 错误

功能说明

- UnreadRune取消上次ReadRune读取的Unicode字符；如果上次读取操作不是ReadRune，则返回错误。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewReader([]byte("你好世界"))
		var buff [3]byte
		// 用Read读取第一个Unicode字符'你'
		b.Read(buff[:])
		fmt.Println(b.UnreadRune())
		// 用ReadRuen读取第二个Unicode字符'好'
		r, _, _ := b.ReadRune()
		fmt.Println(string(r))
		// 取消读取字符'好'
		fmt.Println(b.UnreadRune())
		// 重新读取字符'好'
		r, _, _ = b.ReadRune()
		fmt.Println(string(r))
	}

代码输出
	
	bytes.Reader: previous operation was not ReadRune
	好
	<nil>
	好
