##func Error(w ResponseWriter, error string, code int) 

参数

- w 响应的写入器
- error 用户定义的错误信息
- code  用户定义的http代码

返回值

- 无

函数功能 

- 该函数通过一个ResponseWriter对象输出指定的错误信息和HTTP代码。

例子

  package main
	
	import (
		_ "io"
		"log"
		"net/http"
	)
	
	// hello world, the web server
	func HelloServer(w http.ResponseWriter, req *http.Request) {
		// 正常情况的输出,暂时注释掉了
		//io.WriteString(w, "hello, world!\n")
	
		// 输出用户自定义错误
		http.Error(w, "this is a error", 404)
	}
	
	func main() {
		// 指定当用户访问 http://www.xxx.com:mmmm/hello 的时候(注意，请不要在hello后面加上/变成hello/)
		// 调用HelloServer这个函数来处理
		http.HandleFunc("/hello", HelloServer)
	
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



