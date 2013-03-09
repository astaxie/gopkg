## func TokenString(tok rune) string

参数列表：

- tok 待转换字符

返回值：

- 返回可打印字符串

功能说明：

这个函数用来将tok转换为字符串

代码实例：

	package main

	import(
		"fmt"
		"text/scanner"
	)
	func main(){

		toks := []rune("abcdef")

		for _, tok := range toks{
			fmt.Print(scanner.TokenString(tok))
		}
	}

