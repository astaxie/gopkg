## func ListenIP(netProto string, laddr *IPAddr) (*IPConn, error)

参数列表:

- netProto 网络类型 字符串
- laddr 本地地址指针

返回列表:

- IPConn ip连接指针
- error 错误信息

监听netProto协议和本地地址,返回IP连接指针

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		ipaddr,_ := net.ResolveIPAddr("tcp","0.0.0.0")
		ipconn,err := net.ListenIP("ip:1",ipaddr)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ipconn:%#v",ipconn) //返回 ipconn:&net.IPConn{fd:(*net.netFD)(0xf840069000)}
	}