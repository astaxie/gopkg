## func LookupHost(host string) (addrs []string, err error)

参数列表:

- host 服务器字符串

返回列表:

- 地址列表
- err 错误信息

返回该域名的所有地址字符串列表

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		addrs,err := net.LookupHost("www.163.com")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(addrs) // 打印结果应该类似 [122.228.218.179 183.129.198.96]
	}