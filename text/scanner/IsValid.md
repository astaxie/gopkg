## func (pos *Position) IsValid() bool

参数列表：

- 无

返回值：

- 当期位置是否合法，bool类型

功能说明：

如果当期位置合法，返回true；如果当前位置不合法，返回false。

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

		s.Next()

		pos := s.Pos()
		fmt.Println(pos.IsValid())

	}

