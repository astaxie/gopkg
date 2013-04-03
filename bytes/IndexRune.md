## func IndexRune(s []byte, r rune) int

参数列表

- s 字节切片
- r Unicode字符

返回值

- int s中的位置索引（从0开始）

功能说明

- IndexRune把s解释为UTF-8字节序列，并返回r在s中的位置索引；如果s中不包含r则返回-1

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("你好世界")
		index := bytes.IndexRune(s, '好')
		if index >= 0 {
			fmt.Println(index)
			fmt.Println(string(s[index:]))
		}
	}

代码输出
	
	3
	好世界

