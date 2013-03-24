## func StatusText(code int) string 

参数列表

- code 状态码

返回值：

- string 对应状态码的字符串

功能说明：
该函数为一个http状态码返回对应的文本。
如果状态码未知的话，则返回一个空的字符串。

代码实例

  package main
	
	import (
		"fmt"
		"net/http"
	)
	
	func main() {
		fmt.Println(http.StatusText(404))
		fmt.Println(http.StatusText(202))
	}



