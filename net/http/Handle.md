##func Handle(pattern string, handler Handler) 

参数

- pattern 需要匹配的pattern，比如/hello, /index.php , /index.aspx等等
- handler 实际响应该pattern的函数


返回值

- 无

函数功能 

- 该函数在缺省的服务器路由（DefaultServeMux）中注册一个函数。当一个请求request进来的时候，缺省的服务器路由会依次根据ServeMux.m中的string（路由表达式）来一个一个匹配。文档（服务器路由）ServeMux解释了路由表达式是如何被匹配的。

例子

  package main
	
	import (
		"io"
		"log"
		"net/http"
	)
	
	// hello world, the web server
	func HelloServer(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello, world!\n")
	}
	
	func IndexPhpServer(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<html><body>hello,this is a go page for index.php</body></html>\n")
	}
	
	func main() {
		// 指定当用户访问 http://www.xxx.com:mmmm/hello 的时候(注意，请不要在hello后面加上/变成hello/)
		// 调用HelloServer这个函数来处理
		http.HandleFunc("/hello", HelloServer)
	
		// 调用 http://localhost:8888/index.php 访问
		http.HandleFunc("/index.php", IndexPhpServer)
	
		// 侦听本地的8888端口
		// 客户可以通过浏览器来访问
		// 可以输入 http://localhost:8888/hello来访问
		// 其中localhost会被转换为127.0.0.1,这是本地ip地址
		// 这里需要注意的问题有
		// 1. 360误报，可以删除360装qq管家
		// 2. 某些浏览器不识别这个端口，需要手动配置一下,或者你可以使用80端口，8080端口
		//    或者换成chrome浏览器尝试一下
		// 3. 本地防火墙阻止
		// 如果顺利的话，你的浏览器会输出 this is a error
	
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}




