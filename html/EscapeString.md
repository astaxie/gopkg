#func EscapeString(s string) string

##参数
- s string 需要转义的字符串

##返回
- string 转义后的字符串

##功能说明
EscapeString用于将特殊字符转移为html实体,如把`<`转义成`&lt;` 	
它只会转义下列五种字符: `<` `>` `&` `'`  `"`		
需要注意的是 `UnescapeString(EscapeString(s)) == s` 返回结果一定是true,但是反过来就不一定是true了

##代码示例

	package main
	
	import "fmt"
	import "html"
	
	func main() {
	    var s string = "<script>alert('xss')</script>"
	    fmt.Println(html.EscapeString(s))
	}

输出结果:
	`&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;`