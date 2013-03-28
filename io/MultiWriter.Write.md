# func (t *multiWriter) Write(p []byte) (n int, err error)

参数：
- p 写入的数据

返回值：
- n 写入字节数
- err 写入是否成功，nil代表成功

功能说明：
- 将p中的数据同时写入MultiWriter的多个Writer中

示例：
  package main
	
	import (
		"io"
		"fmt"
		"os"
	)
	
	func main() {
		writer1, _ := os.Create("dst1.txt")
		writer2, _ := os.Create("dst2.txt")
		multiWriter := io.MultiWriter(writer1, writer2)
		p := []byte("Write to multi dest")
		n, err := multiWriter.Write(p)
		if err == nil {
			fmt.Println("Write success total", n)
		}
	}
