## func ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error)

参数列表

- fixedURL 代理的URL

返回值：

- func(*Request) (*url.URL, error) 返回一个代理函数

功能说明：
ProxyURL返回一个代理函数，这个代理函数（一般在Transport中使用）接受请求，并总是返回一个（代理后的）地址。


代码实例

  package main
	
	import (
		"fmt"
	)
	
	func main() {
		fmt.Println("这个函数暂时不知道用在什么地方")
	}


