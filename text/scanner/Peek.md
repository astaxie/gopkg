## func (s *Scanner) Peek() rune

参数列表：

- 无

返回值：

- 返回下一个字符，rune类型

功能说明：

在不使词法解析器向前进一位的情况下，返回下一个字符。如果解析器当前位置已经在源的最后一位，返回EOF。

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

		//this will print "i" to stdout
		fmt.Println(scanner.TokenString(s.Peek()))
	}

