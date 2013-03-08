## func LookupAddr(addr string) (name []string, err error)

参数列表
- addr 地址字符串

返回列表
- name 主机名列表
- err 错误信息


返回系统网络设备信息列表

实例:
	package main
	
	import "fmt"
	import "net"
	
	func main() {
		addr := net.JoinHostPort("gocn_server","8080")
		fmt.Println(addr) // 打印结果应该为 gocn_server:8080
	}