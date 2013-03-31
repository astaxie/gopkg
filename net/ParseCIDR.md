## func ParseCIDR(s string) (IP, *IPNet, error)

参数列表:

- s 字符串

返回列表:

- IP IP结构
- IPNet 结构
- error 错误信息

解析 CIDR(无类域间路由) 格式字符串,返回 IP结构,IPNet结构和错误信息

[百度百科 - CIDR](http://baike.baidu.com/view/4217886.htm)

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		ip,ipnet,err := net.ParseCIDR("198.168.0.0/16")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("IP:%#v",ip) //返回 IP:[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xc6, 0xa8, 0x0, 0x0}
		fmt.Printf("IPNet:%#v",ipnet) //返回 IPNet:&net.IPNet{IP:[]byte{0xc6, 0xa8, 0x0, 0x0}, Mask:[]byte{0xff, 0xff, 0x0, 0x0}}
	}