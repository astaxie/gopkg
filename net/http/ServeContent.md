## func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker) 

参数列表

- w 服务器响应
- req 客户端请求
- name 文件的名称
- modtime 文件的修改时间 
- content 文件的内容，必须实现io.ReadSeeker这个接口中的方法

返回值：

- error 成功为nil

功能说明：

ServeContent使用ReadSeeker所读取的内容回复给用户请求.
ServeContent比io.Copy更好的是，他能够合适的处理一批请求，设置MIME类型，并且能够处理文件是否修改的请求。
如果响应的内容类型头没有设置,该函数首先会尝试从文件的文件扩展名推断文件类型。
如果推断不出来，则会读取文件的第一个块并传送给DetectContentType来检测类型。
文件名称也可以不使用。
如果文字名称为空，则服务器不会传送给响应。
如果修改时间不为0，ServeContent会把它放在服务器响应的Last-Modified头里面。
如果客户端请求中包含了If-Modified-Since头，ServeContent会使用modtime来判断是否把内容传给客户端。


content的Seek方法必须能够工作。
ServeContent通过定位到文件结尾来确定文件大小。
*os.File中实现了io.ReadSeeker接口。

代码实例

  package main
	
	import (
		"fmt"
		"log"
		"net/http"
		"time"
	)
	
	// 实现一个byte的ReadSeeker
	type MySeeker []byte
	
	// 仅仅是copy输出
	func (d MySeeker) Read(p []byte) (n int, err error) {
		n = copy(p, d)
		fmt.Printf("read is", n)
		return n, nil
	}
	func (d MySeeker) Seek(offset int64, whence int) (ret int64, err error) {
		ret = int64(len(d))
		fmt.Printf("ret is", ret)
		return ret, nil
	}
	
	func Handler(w http.ResponseWriter, r *http.Request) {
		// 仅仅向客户端输出几个字节
		var seeker MySeeker = []byte{'H', 'E', 'L', 'L', 'O', 0}
		
		// 文件名称随便填写了1个
		http.ServeContent(w, r, "1.txt", time.Now(), seeker)
	}
	func main() {
		http.HandleFunc("/hello", Handler)
	
		s := &http.Server{
			Addr: ":8888",
		}
		log.Fatal(s.ListenAndServe())
	}
