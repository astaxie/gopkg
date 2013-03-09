## func (pos Position) String() string

参数列表：

- 无

返回值：

- 当前位置，string类型

功能说明：

返回当前位置的字符串表示，格式为："行：列"

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
	//this will print "1:4" to stdout
	fmt.Println(s.Pos().String())

	s.Next()
	//this will print "1:5" to stdout
	fmt.Println(s.Pos().String())
}

