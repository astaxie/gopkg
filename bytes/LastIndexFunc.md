## func LastIndexFunc(s []byte, f func(r rune) bool) int

参数列表

- s 字节切片
- f 检查函数

返回值

- int s中字节的位置索引

功能说明

- LastIndexFunc把s解释为UTF-8编码的字节序列，返回满足f(c)==true的字符c在s中最后一次出现的位置索引；如果没有找到这样的字符则返回-1。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		f1 := func(r rune) bool { return r == '7' }
		f2 := func(r rune) bool { return r == '9' }
		s := []byte("12345678")
		fmt.Println(bytes.LastIndexFunc(s, f1))
		fmt.Println(bytes.LastIndexFunc(s, f2))
	}

代码输出

	6
	-1
