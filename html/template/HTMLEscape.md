#func HTMLEscape(w io.Writer, b []byte)

##参数
- w io.Writer 	写入的Writer
- b []byte 		需要进行转义的内容

##功能说明
HTMLEscape用于将b转义后的HTML文本写入到w中.

##代码示例

	package main
	
	import (
		"bytes"
		"fmt"
		"html/template"
	)
	
	func main() {
	
		w := bytes.NewBufferString("")
	
		b := []byte("<script>alert('xss!')</script>")
	
		template.HTMLEscape(w, b)
	
		fmt.Println(w)
	
	}
	
输出:	
	`&lt;script&gt;alert(&#39;xss!&#39;)&lt;/script&gt;`
