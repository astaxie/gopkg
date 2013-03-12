#func QueryEscape(s string) string

[参数列表]：

- s 表示需要参数(字符串) 

[返回值]：

-返回值为：转义字符串

[功能说明]：

escapes the string so it can be safely placed inside a URL query.

逃避查询，因此它可以被安全的放置在URL查询字符串。请参阅RFC 3986。

[代码实例]：

	package main	
	import (
		"fmt"
		"net/url"
	)
	
	func main() {
		sUrl := url.QueryEscape("http://www.golang.org")
		fmt.Println(sUrl)
	}
	//输出结果：http%3A%2F%2Fwww.golang.org

