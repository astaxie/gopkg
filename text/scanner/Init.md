## func (s *Scanner) Init(src io.Reader) *Scanner

参数列表：

- src 输入源

返回值：

- 初始化的词法解析器

功能说明：

初始化一个词法解析器。

代码实例：

	package main

	import(
		"fmt"
		"strings"
		"text/scanner"
	)

	func main(){

		src := strings.NewReader("int hello = 3; hello+23; print hello;")

		fmt.Println(src)

		var s scanner.Scanner
		s.Init(src)

		tok := s.Scan()
		fmt.Println(s.TokenText())
		for tok != scanner.EOF{
			tok = s.Scan()
		fmt.Println(s.TokenText())
		}
	}

