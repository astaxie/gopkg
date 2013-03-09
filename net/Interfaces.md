## func Interfaces() ([]Interface, error)

参数列表
- 无

返回列表
- []Interface 系统网络设备信息列表
- error 错误信息

返回系统网络设备信息列表

实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		eths,err := net.Interfaces()
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(eths) // 打印结果应该类似[{1 16436 lo  up|loopback} {2 1500 eth0 00:16:3e:01:04:1f up|broadcast|multicast} {3 1500 eth1 00:16:3e:01:04:1e up|broadcast|multicast}
	}