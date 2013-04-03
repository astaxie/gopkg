## func LookupPort(network, service string) (port int, err error)

参数列表:

- network 网络字符串
- service 服务名

返回列表:

- port 端口号
- err 错误信息

查找特定网络的端口和服务

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		port,err := net.LookupPort("tcp","https")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(port) //返回 443
	}