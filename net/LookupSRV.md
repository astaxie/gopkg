## func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)

参数列表:

- service 服务名
- proto 协议名
- name 域名

返回列表:

- cname cname字符串
- addrs SRV记录列表
- err 错误信息

查找该服务和协议下的SRV记录地址

[百度百科-SRV记录](http://baike.baidu.com/view/1372993.htm)

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		cname,addrs,err:=net.LookupSRV("xmpp-server","tcp","gocn.im")
		if err != nil {
			fmt.Println(err)
		}
		for _,addr := range addrs {
			fmt.Println(addr) // 返回类似 &{xmpp-server.l.google.com. 5269 5 0}
		}
		fmt.Println(cname) //返回类似 _xmpp-server._tcp.gocn.im.
	}