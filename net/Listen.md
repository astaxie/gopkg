## func Listen(net, laddr string) (Listener, error)

参数列表:

- net 网络字符串
- laddr 本地地址字符串

返回列表:

- Listener 监听者
- err 错误信息

查找特定网络的端口和服务

代码实例:

	package main
	
	import "io"
	import "log"
	import "net"
	
	func main() {
		l, err := net.Listen("tcp", ":2000")
		if err != nil {
		    log.Fatal(err)
		}
		for {
		    // 等待下一个连接,如果没有连接,l.Accept会阻塞
		    conn, err := l.Accept()
		    if err != nil {
		        log.Fatal(err)
		    }
		    // 将新连接放入一个goroute里,然后再等下一个新连接.
		    go func(c net.Conn) {
		        // 显示连接发送来的内容
		        io.Copy(c, c)
		        // 关闭连接
		        c.Close()
		    }(conn)
		}
	}