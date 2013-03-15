## func IndexFunc(s []byte, f func(r rune) bool) int

参数列表

- s 要检查的字节切片
- f 过滤函数

返回值

- int s中的位置索引（从0开始）

功能说明

- IndexFunc把s解释成UTF-8字节序列，并返回第一个满足f(c)==true的字符c的位置索引；如果没有字符满足f(c)==true则返回-1。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s := []byte("你好世界")
		fmt.Println(bytes.IndexFunc(s, func(r rune) bool {
			return r == '好'
		}))
		fmt.Println(bytes.IndexFunc(s, func(r rune) bool {
			return r == '!'
		}))
	}

代码输出

	3
	-1
