## func LookupAddr(addr string) (name []string, err error)

参数列表

- addr  地址字符串

返回列表

- name  主机名列表
- err   错误信息

功能说明

- 通过IP地址查询域名列表

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		names := net.LookupAddr("173.194.127.81")
		fmt.Println(addr)
	}

代码输出：

    [hkg03s11-in-f17.1e100.net.]
    <nil>