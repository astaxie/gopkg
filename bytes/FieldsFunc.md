## func FieldsFunc(s []byte, f func(rune) bool) [][]byte

参数列表

- s 字节切片
- f 过滤函数

返回值

- [][]byte 分隔的字节切片的切片

功能说明

- FieldsFunc把s解释为UTF-8编码的字符序列，对于每个Unicode字符c，如果f(c)返回false就把c作为分隔字符对s进行拆分。如果没有字符满足f(c)为true，则返回空的切片。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		fields1 := bytes.FieldsFunc([]byte("ab,cd,e"), func(c rune) bool {
			return c == rune(',')
		})
		fields2 := bytes.FieldsFunc([]byte("你好啊世界"), func(c rune) bool {
			return c == rune('啊')
		})
		for i, s := range fields1 {
			fmt.Printf("%d: %s\n", i, string(s))
		}
		for i, s := range fields2 {
			fmt.Printf("%d: %s\n", i, string(s))
		}
	}

代码输出

	0: ab
	1: cd
	2: e
	0: 你好
	1: 世界
