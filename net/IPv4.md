## func IPv4(a, b, c, d byte) IP

参数列表:

- a IP地址的第1个字节
- b IP地址的第2个字节
- c IP地址的第3个字节
- d IP地址的第4个字节

返回列表:

- IP ip结构

函数功能:

- 根据a,b,c,d四个字节返回ip结构

代码实例:

    package main
    
    import "fmt"
    import "net"
    
    func main() {
        ip := net.IPv4(127,0,1,1)
        fmt.Println(ip)
    }

代码输出:

    127.0.1.1
