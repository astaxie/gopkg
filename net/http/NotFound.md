## func NotFound(w ResponseWriter, r *Request) 

参数

- w 响应的写入器
- r 用户请求

返回值

- 无

函数功能 

- 该函数通过一个ResponseWriter对象输出404 page not found。

例子
  package main
	
	import (
		_ "io"
		"log"
		"net/http"
	)
	
	func HelloServer(w http.ResponseWriter, req *http.Request) {
		http.NotFound(w, req)
	}
	
	func main() {
	
		http.HandleFunc("/hello", HelloServer)
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
