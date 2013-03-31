#func HTMLEscapeString(s string) string

##参数
- s string 需要进行转义的字符串

##返回
- string 转义后的字符串

##功能说明
HTMLEscapeString返回s转义后的HTML文本.

##代码示例

	package main
	
	import (
		"fmt"
		"html/template"
	)
	
	func main() {
	
		s := "<script>alert('xss!')</script>"
	
		fmt.Println(template.HTMLEscapeString(s))
	
	}
	
输出:	
	`&lt;script&gt;alert(&#39;xss!&#39;)&lt;/script&gt;`
