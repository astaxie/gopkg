## func DialIP(netProto string, laddr, raddr \*IPAddr) (\*IPConn, error)

参数列表:

- netProto 网络类型协议 字符串
- laddr 本地地址
- raddr 远端地址

netProto必须为以下三种之一

- "ip"
- "ip4"
- "ip6"

返回列表:

- IPConn ip连接指针
- error 错误信息

用连接网络类型协议和本地地址连接远端地址,返回IP连接

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		raddr, _ := net.ResolveIPAddr("tcp", "www.163.com")
		ipconn, err := net.DialIP("ip:1", nil, raddr)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ipconn:%#v", ipconn) //ipconn:&net.IPConn{fd:(*net.netFD)(0xf840069000)}
	}
