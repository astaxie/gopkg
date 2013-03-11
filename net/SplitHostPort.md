## func SplitHostPort(hostport string) (host, port string, err error)

参数列表:

- hostport 主机端口字符串

返回列表:

- host 主机字符串
- port 端口字符串
- err 错误信息

分解主机端口字符串


代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		host,port,err:=net.SplitHostPort("163.com:80")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(host) // 返回 163.com
		fmt.Println(port) // 返回 80
		
		host1,port1,err1:=net.SplitHostPort("[abc:bbc]:88")
		if err1 != nil {
			fmt.Println(err1)
		}
		fmt.Println(host1) // 返回 abc:bbc
		fmt.Println(port1) // 返回 88
	}