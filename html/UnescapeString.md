#func UnescapeString(s string) string

##参数
- s string 需要反转义的字符串

##返回
- string 反转义后的字符串

##功能说明
UnescapeString用于将html实体转义回特殊字符,如把`&lt;`转义成`<` 	
此函数可以反转义的html实体比EscapeString可转义的更多,EscapeString只会转义下列五种字符: `<` `>` `&` `'`  `"`	 	
例如,`&aacute;` `&#225;` `&nbsp;` 这些都可以被反转义为特殊字符	
需要注意的是 `UnescapeString(EscapeString(s)) == s` 返回结果一定是true,但是反过来就不一定是true了

##代码示例

	package main
	
	import "fmt"
	import "html"
	
	func main() {
		var s string = "&#225;&nbsp;&#225;&nbsp;&#225;"
		fmt.Println(html.UnescapeString(s))
	}
	
输出结果:		
	`á á á`