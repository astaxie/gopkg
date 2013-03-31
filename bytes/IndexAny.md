## func IndexAny(s []byte, chars string) int

参数列表

- s 要检查的字节切片
- chars 要检查是否包含字符的字符串

返回值

- int s中的位置索引（从0开始）

功能说明

- IndexAny把s解释为UTF-8编码的字节序列，返回chars中任一个字符在s中第一次出现的位置索引；如果s中不包含chars中任何一个字符则返回-1.

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		s := []byte("大家好大家早")
		index := bytes.IndexAny(s, "好早")
		if index >= 0 {
			fmt.Printf("%d: %s\n", index, string(s[index:]))
		}
	}
	
代码输出

	6: 好大家早
