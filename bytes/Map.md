## func Map(mapping func(r rune) rune, s []byte) []byte

参数列表

- mapping 映射函数
- s 字节切片

返回值

- []byte 被映射后的s的副本字节切片

功能说明

- Map把s解释为UTF-8编码的字节序列，并把s中的每个Unicode字符用mapping函数获得对应的字符并存放到一个新创建的字节切片的对应位置，并返回此新创建的字节切片。

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		s1 := []byte("大家上午好")
		s2 := []byte("12345678")
		m1 := func(r rune) rune {
			if r == '上' {
				return '下'
			}
			return r
		}
		m2 := func(r rune) rune {
			return r + 1
		}
		fmt.Println(string(bytes.Map(m1, s1)))
		fmt.Println(string(s1))
		fmt.Println(string(bytes.Map(m2, s2)))
		fmt.Println(string(s2))
	}

代码输出

	大家下午好
	大家上午好
	23456789
	12345678

	
