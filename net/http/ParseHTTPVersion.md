## func ParseHTTPVersion(vers string) (major, minor int, ok bool) 

参数

- vers 响应的写入器
- r 用户请求

返回值

- major 主版本号
- minor 从版本号
- ok    true或者false

函数功能 

- 该函数解析一个HTTP 版本字符串. 比如输入"HTTP/1.0" 将会返回 (1, 0, true). 

例子
  package main
	
	import (
		"fmt"
		"net/http"
	)
	
	func main() {
		major, minor, ok := http.ParseHTTPVersion("HTTP/1.0")
		fmt.Println(major, minor, ok)
	}

