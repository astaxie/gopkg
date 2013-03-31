## func LookupMX(name string) (mx []*MX, err error)

参数列表:

- name 域名字符串

返回列表:

- mx记录列表
- err 错误信息

[MX记录-百度百科](http://baike.baidu.cn/view/65595.htm)

返回该域名的所有mx记录列表

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		mxs,err := net.LookupMX("163.com")
		if err != nil {
			fmt.Println(err)
		}
		for _,mx := range mxs{
			fmt.Println(mx) 
			// 打印结果应该类似 
			// &{163mx02.mxmail.netease.com. 10}
			// &{163mx01.mxmail.netease.com. 10}
			// &{163mx03.mxmail.netease.com. 10}
			// &{163mx00.mxmail.netease.com. 50}
		}
	}