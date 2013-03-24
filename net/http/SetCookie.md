## func SetCookie(w ResponseWriter, cookie *Cookie) 

参数列表

- w 对客户端的响应
- cookie cookie的内容

返回值：

- 无

功能说明：

该函数添加一个Set-Cookie头到客户端响应中


代码实例

  package main
	
	import (
		"log"
		"net"
		"net/http"
		"time"
	)
	
	func HelloServer(w http.ResponseWriter, req *http.Request) {
		expire := time.Now().AddDate(0, 0, 1)
		mycookie := http.Cookie{"test", "testcookie", "/", "www.sanguohelp.com", expire, expire.Format(time.UnixDate),
			86400, true, true, "test=testcookie", []string{"test=tcookie"}}
	
		http.SetCookie(w, &mycookie)
	
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



