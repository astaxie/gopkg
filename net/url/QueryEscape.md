#func QueryEscape(s string) string

参数列表：

- s 表示需要参数(字符串) 

返回值：
-返回值为：转义字符串

功能说明：
这个函数主要是用来将s中包裝的網址转义成可进可查询的字符串

代码实例：

	package main	
	import (
		"fmt"
		"net/url"
	)
	
	func main() {
		var sUrl = url.QueryEscape("http://www.163.com")
		fmt.Printf(sUrl)
	}
	//输出结果：http%A(MISSING)%F(MISSING)%F(MISSING)www.163.com

