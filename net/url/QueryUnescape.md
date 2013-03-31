#func QueryUnescape(s string) (string, error)

[参数列表]：

- s 表示需要参数(字符串) 

[返回值]：

-返回值为：字符串,错误信息

[功能说明]：
对QueryEscape函数转换的值进行逆转。如果%后没有包含两个十六进制数，它将返回一个错语

[代码实例]：

  	package main	
	import (
		"fmt"
		"net/url"
	)
	
	func main() {
		uUrl,err := url.QueryUnescape("http%3A%2F%2Fwww.golang.org")
		if err == nil {
			fmt.Printf(uUrl)
		}else{
			fmt.Println(err)
		}
	}
	//输出结果：http://www.golang.org

