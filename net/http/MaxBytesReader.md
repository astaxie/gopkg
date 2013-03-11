##func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser 

参数

- w ResponseWriter
- r io.ReadCloser，一般可以指向req.Body
- n 限制的大小

返回值

- io.ReadCloser 

函数功能 

- MaxBytesReader跟io.LimitReader函数很像。但是它被设计来设置接收的请求体的最大大小。
跟io.LimitReader不同MaxBytesReader的返回值是一个ReadCloser，当读取超过限制时会返回non-EOF错误。
并且当它调用关闭方法的时候会把潜在的读取者（函数/进程）也关闭掉。

MaxBytesReader用来保护服务器端，以避免客户端偶然或者恶意发送的长数据请求导致的server资源的浪费。 


例子 服务器端
调用这个函数限制客户端上传数据为10个字节

  package main
	
	import (
		"fmt"
		"io"
		"io/ioutil"
		"log"
		"net/http"
	)
	
	// hello world, the web server
	func HelloServer(w http.ResponseWriter, req *http.Request) {
		fmt.Println("come on")
	
		req.Body = http.MaxBytesReader(w, req.Body, 10)
		n, err := io.Copy(ioutil.Discard, req.Body)
		if err == nil {
			fmt.Println("iocpy err")
		}
		if n != 20 {
			fmt.Println("io.Copy = ", n, ", want 20")
		}
	
		io.WriteString(w, "hello, world!hello, world!hello, world!hello, world!hello, world!\n")
	
	}
	
	func main() {
	
		// 指定当用户访问 http://www.xxx.com:mmmm/hello 的时候(注意，请不要在hello后面加上/变成hello/)
		// 调用HelloServer这个函数来处理
		http.HandleFunc("/hello", HelloServer)
	
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	
	}


客户端发送JSON数据

	package main
	
	import (
		"fmt"
		"net/http"
		"strings"
	)
	
	func main() {
		json := `{"content":"hello,world"}`
		b := strings.NewReader(json)
	
		http.Post("http://localhost:8888/hello", "image/jpeg", b)
		fmt.Println("post ok")
	}




