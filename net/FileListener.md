## func FileListener(f *os.File) (l Listener, err error)

参数列表:

- f 文件名

返回列表:

- l 监听者
- err 错误

打开一个文件,并负责他的关闭,返回一个监听者或者错误

代码实例:

	package main
	
	import "fmt"
	import "net"
	
	func main() {
		port,err := net.LookupPort("tcp","https")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(port) //返回 443
	}