## func LookupCNAME(name string) (cname string, err error)

参数列表:

- name 域名字符串

返回列表:

- 别名字符串
- error 错误信息

返回该域名的CNAME记录

[CNAME-百度百科](http://baike.baidu.com/view/552919.htm)

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		cname,err := net. LookupCNAME("bbs.gocn.im")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cname) // 打印结果应该为 bbs.gocn.im.
	}