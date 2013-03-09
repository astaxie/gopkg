## func (s *Scanner) TokenText() string

参数列表：

- 无

返回值：

- 当前被扫描到的token，string类型

功能说明：

返回当前被扫描到token的字符串表示。在调用Scan之后使用

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

		//this will print nothing
		fmt.Print(s.TokenText())

		s.Scan()

		//this will print "int" to stdout
		fmt.Print(s.TokenText())

	}


