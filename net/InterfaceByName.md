## func InterfaceByName(name string) (*Interface, error)

参数列表:

- name 网络接口名称

返回列表:

- I*nterface 接口结构指针
- error 错误信息

根据索引查找网络接口

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		interf,err := net.InterfaceByName("eth0")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(interf) //返回 &{2 1500 eth0 00:16:3e:01:04:1f up|broadcast|multicast}
	}