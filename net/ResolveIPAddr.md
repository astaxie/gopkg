## func ResolveIPAddr(net, addr string) (*IPAddr, error)

参数列表:

- net 网络类型 字符串
- addr 地址 字符串

返回列表:

- IPAddr ip地址指针
- error 错误信息

返回网络类型下的ip地址/域名的IPAddr结构指针

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		ipaddr,err := net.ResolveIPAddr("tcp","www.163.com")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ipaddr:%#v",ipaddr) //ipaddr:&net.IPAddr{IP:[]byte{0x7a, 0xe4, 0xda, 0xb3}}
	}