## func (s *Scanner) Next() rune

参数列表：

- 无

返回值：

- 返回rune

功能说明：

读取并返回下一个字符。如果读到源的结尾，返回EOF。如果s.Error不为nil，表示遇到了错误；并输出错误信息到控制台。Next不会更新词法解析器的Position值；可以使用Pos()获取当前位置。

代码实例：

	package main

	import (
		"fmt"
		"strings"
		"text/scanner"
	)

	func main() {

		src := strings.NewReader("int num = 1;")

		var s scanner.Scanner

		s.Init(src)

		//this will print "i" to stdout
		fmt.Println(scanner.TokenString(s.Next()))

	}

