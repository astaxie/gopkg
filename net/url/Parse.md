#func Parse(rawurl string) (url *URL, err error)
[参数列表]：

- rawurl 网址

[返回值]：

- url rawurl的结构
- err 错误信息

[功能说明]：
解析一个rawurl的结构。 rawurl可以是相对的或绝对的。
[代码实例]：

    package main
    
    import (
      "fmt"
    	"log"
    	"net/url"
    )
    
    func main() {
    	u, err := url.Parse("http://www.bing.com/search?q=dotnet")
    	if err != nil {
    		log.Fatal(err)
    	}
    	fmt.Println(u.Scheme)
    	//输出结果为：http
    	fmt.Println(u.Host)
    	//输出结果为：www.bing.com
    	fmt.Println(u.Path)
    	//输出结果为：/search
    	fmt.Println(u.RawQuery)
    	//输出结果为：/search
    	fmt.Println(u.Opaque)
    }
