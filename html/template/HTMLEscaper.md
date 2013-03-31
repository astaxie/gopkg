#func HTMLEscaper(args ...interface{}) string

##参数
- args 多个输入参数

##返回
- string 转义后的字符串

##功能说明
HTMLEscaper 返回多个参数一起转义后的字符串

##代码示例

	package main
	
	import (
		"fmt"
		"html/template"
	)
	
	func main() {
	
		a := "<script>alert('xss!')</script>"
	
		b := "<h1>Raymond</h1>"
	
		c := "<h2>Chou</h2>"
	
		fmt.Println(template.HTMLEscaper(a, b, c))
	
	}
	
输出:	
`&lt;script&gt;alert(&#39;xss!&#39;)&lt;/script&gt;&lt;h1&gt;Raymond&lt;/h1&gt;&lt;h2&gt;Chou&lt;/h2&gt;`
