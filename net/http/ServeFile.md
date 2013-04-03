## func ServeFile(w ResponseWriter, r *Request, name string) 

参数列表

- w 响应
- r 客户端请求
- name 文件名或者文件目录名称

返回值：

- 无

功能说明：
该函数实现了向客户请求输出目录或者文件内容。

代码实例

  package main
	
	import (
		"log"
		"net"
		"net/http"
	)
	
	func HelloServer(w http.ResponseWriter, req *http.Request) {
	
		http.ServeFile(w, req, "index.html")
	
	}
	
	func main() {
	
		http.HandleFunc("/hello", HelloServer)
	
		l, e := net.Listen("tcp", ":8888")
		if e != nil {
			log.Fatal("Listen: ", e)
		}
	
		err := http.Serve(l, nil)
		if err != nil {
			log.Fatal("Serve: ", err)
		}
	}

