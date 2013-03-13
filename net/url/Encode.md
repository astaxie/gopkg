#func (Values) Encode
[参数列表]：

- 

[返回值]：

- 

[功能说明]：
转为URL编码形式。例如"foo=bar&bar=baz"

[代码实例]：
    package main
    
    import (
      "fmt"
    	"net/url"
    )
    func main() {
    	v := url.Values{}
    	v.Add("foo", "bar")
    	v.Add("bar", "baz")
    
    	fmt.Println(v.Encode())
    	//输出结果：foo=bar&bar=baz
    }
