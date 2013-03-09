## func (s *Scanner) Scan() rune

参数列表：

- 无

返回值：

- 下一个字符或token

功能说明：

Scan方法读取源中下一个token或字符，并返回。如果读到源的结尾，返回EOF。

代码实例：

	package main

	import(
		"fmt"
		"strings"
		"text/scanner"
	)

	func main(){
	
		src := strings.NewReader("int num = 1;")
		var s scanner.Scanner

		s.Init(src)

		s.Scan()
		//this will print the next token "int "to stdout
		fmt.Println(s.TokenText())
	}

