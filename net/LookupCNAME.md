## func LookupCNAME(name string) (cname string, err error)

参数列表:

- name 域名字符串

返回列表:

- cname 主机名
- error 错误信息

返回该域名的CNAME记录

[CNAME-百度百科](http://baike.baidu.com/view/552919.htm)

代码实例:

	package main
	
	import "fmt"
	import "net"
	
    func main() {
        fmt.Println( net. LookupCNAME("any.weizhe.net"))
        fmt.Println( net. LookupCNAME("blog.weizhe.net"))
    }
    
代码输出：

    blog.weizhe.net. <nil>
    blog.weizhe.net. <nil>