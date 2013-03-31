## func Dial(net, addr string) (Conn, error)

参数列表:

- net 网络类型 字符串
- addr 地址信息 字符串

返回列表:

- Conn 连接接口
- err 错误信息

使用相应的网络类型连接远端地址,并返回连接接口或错误信息

网络类型可以是以下几种

- "tcp"
- "tcp4" (仅限IPv4)
- "tcp6" (仅限IPv6)
- "udp"
- "udp4" (仅限IPv4)
- "udp6" (仅限IPv6)
- "ip"
- "ip4" (仅限IPv4)
- "ip6" (仅限IPv6)
- "unix"
- "unixpacket"

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		conn,err := net.Dial("tcp","www.163.com:80")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%#v",conn) //返回类似 &net.TCPConn{fd:(*net.netFD)(0xf840069000)}
	}