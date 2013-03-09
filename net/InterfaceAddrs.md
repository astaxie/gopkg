## func InterfaceAddrs() ([]Addr, error)

参数列表:

- 无

返回列表:

- []Addr 网络地址列表
- error 错误信息

返回系统网络地址列表,和错误.

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		addrs,err := net.InterfaceAddrs()
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(addrs) // 打印结果应该类似 [127.0.0.1/8 10.200.54.178/20 42.121.14.243/22]
	}
