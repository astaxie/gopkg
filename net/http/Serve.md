## func Serve(l net.Listener, handler Handler) error

参数列表

- l 监听者
- handler Handler参数一般是nil，此时使用DefaultServeMux。

返回值：

- error 成功为nil

功能说明：
Serve接受Listener l接收到的HTTP连接，并为每个连接创建一个新的线程。服务线程会读取每一个请求，调用handler做出回应。Handler参数一般是nil，此时使用DefaultServeMux。

代码实例

  package main
	
	import (
		"io"
		"log"
		"net"
		"net/http"
	)
	
	func HelloServer(w http.ResponseWriter, req *http.Request) {
	
		io.WriteString(w, "hellWorld server")
	
	}
	
	func main() {
	
		http.HandleFunc("/hello", HelloServer)
	
		// 首先，创建用tcp协议监听8888端口
		l, e := net.Listen("tcp", ":8888")
		if e != nil {
			log.Fatal("Listen: ", e)
		}
	
		// 然后在监听的这个端口上启用http服务进行http服务
		err := http.Serve(l, nil)
		if err != nil {
			log.Fatal("Serve: ", err)
		}
	}




