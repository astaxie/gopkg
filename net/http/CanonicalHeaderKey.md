## func CanonicalHeaderKey(s string) string

参数列表

- s 需要做标准化的http header字符串 


返回值：

- 返回按照http协议定义的标准化字符串

功能说明：
这个函数返回一个规范好的标准的http header字符串。
标准化的http header 字符串格式为：
第一个字母和跟着“-”字符后面的第一个字母大写，剩下的字符全部小写。
举个例子：
对于一个http头（header） "accept-encoding" 来说，
规范好的标准化格式就是
"Accept-Encoding"。



代码实例：

package main

import (
  "fmt"
	"net/http"
)

func main() {
	fmt.Println(http.CanonicalHeaderKey("accept-encoding"))
	//Accept-Encoding
	fmt.Println(http.CanonicalHeaderKey("accept-Language"))
	//Accept-Language
}

