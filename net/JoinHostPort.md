## func JoinHostPort(host, port string) string

参数列表
- host string 主机名
- port string 端口号

返回列表
- string 主机名与端口号结合后的字符串


返回主机名与端口号结合后的字符串,类似host:port
比如:
	
	host 为 gocn_server
	port 为 8080
	则返回字符串是 gocn_server:8080
	
如果主机名中存在冒号
比如:

	host 为 gocn:server
	port 为 8080
	则返回的字符串是 [gocn:server]:8080

实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		addr1 := net.JoinHostPort("gocn_server","8080")
		addr2 := net.JoinHostPort("gocn:server","8080")
		fmt.Println(addr1) // 打印结果应该为 gocn_server:8080
		fmt.Println(addr2) // 打印结果应该为 [gocn:server]:8080
	}