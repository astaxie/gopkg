## func ProxyFromEnvironment(req *Request) (*url.URL, error)

参数列表

- req 用户的请求

返回值：

- *url.URL 代理的URL,如果没有使用代理或者代理全局变量没有定义则返回nil
- error 错误，如果没有使用代理或者代理全局变量没有定义则返回nil

功能说明：
ProxyFromEnvironment返回给定request的代理url.
一般该URL由用户的环境变量$HTTP_PROXY和$NO_PROXY (或$http_proxy和$no_proxy)指定。
如果用户的全局代理环境无效则返回一个错误。
如果用户没有使用代理或者全局环境变量没有定义则会返回一个nil的URL和一个nil的错误。


代码实例

  package main
	
	import (
		"io"
		"log"
		"net/http"
	)
	
	func HelloServer(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<html>\n<body>\n")
		io.WriteString(w, "hello, world!<br/>\n")
	
		myurl, _ := http.ProxyFromEnvironment(req)
		if myurl != nil {
			io.WriteString(w, myurl.Scheme+"<br/>\n")
			io.WriteString(w, myurl.Opaque+"<br/>\n")
			io.WriteString(w, myurl.Host+"<br/>\n")
			io.WriteString(w, myurl.Path+"<br/>\n")
			io.WriteString(w, myurl.RawQuery+"<br/>\n")
			io.WriteString(w, myurl.Fragment+"<br/>\n")
		} else {
			io.WriteString(w, "url is null <br/>\n")
		}
		io.WriteString(w, "</body>\n</html>\n")
	
	}
	
	func main() {
	
		http.HandleFunc("/hello", HelloServer)
	
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}



