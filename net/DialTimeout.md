## func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)

参数列表:

- net 网络类型 字符串
- addr 远端地址 字符串
- timeout 超时时间

返回列表:

- Conn 连接接口
- error 错误信息

查找特定网络的端口和服务

代码实例:

	package main

	import "fmt"
	import "net"
	import "time"

	func main() {
		timeOut := time.Duration(10) * time.Second
		conn, err := net.DialTimeout("tcp", "www.163.com:80", timeOut)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%#v", conn) //返回类似 &net.TCPConn{fd:(*net.netFD)(0xf840069000)}
	}