## func Redirect(w ResponseWriter, r *Request, urlStr string, code int) 

参数列表

- w 服务器响应
- r 客户端请求
- urlStr 要重定向的地址
- code 定义在http里面，一般我们可以使用http.StatusFound

返回值：

- 无

功能说明：
Redirect告诉request重定向到一个url,这个URL可以是请求路径的的相对路径。

代码实例

  package main
	
	import (
		"io"
		"log"
		"net/http"
	)
	
	func HelloServer(w http.ResponseWriter, req *http.Request) {
	
		http.Redirect(w, req, "world", http.StatusFound)
	
	}
	
	func WorldServer(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "world server")
	
	}
	
	func main() {
	
		http.HandleFunc("/hello", HelloServer)
		http.HandleFunc("/world", WorldServer)
	
		err := http.ListenAndServe(":9999", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}



