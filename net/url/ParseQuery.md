#func ParseQuery(query string) (m Values, err error)

[参数列表]：

- query 表示需要参数(字符串) 

[返回值]：

-m mps
-err 错误信息

[功能说明]：
通过ParseQuery解析URL编码的查询字符串，并返回一个map ，列出每个键的值
ParseQuery总是返回一个非空的map，包含所有有效的查询参数，通过";"与"&"分割;如果遇到解析异常将通过err返回错误信息

[代码实例]：
  package main
  
  import (
    "fmt"
  	"net/url"
  )
  
  func main() {
  	//http://127.0.0.1:6060/src/pkg/url/url.go?h=%22%26%22;s=14652:14657#L534&test=abc;123
  	test,err := url.ParseQuery("h=%22%26%22;s=14652:14657#L534&test=abc;123")
  	if err == nil {
  		fmt.Println(test)
  	}else{
  		fmt.Println(err)
  	}
  
  }
