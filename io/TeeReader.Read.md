# func (t *teeReader) Read(p []byte) (n int, err error)

参数：
- p 在读取与写入之间缓存数据

返回值：
- n 读取及写入的字节数
- err 读取及写入是否成功；nil代表成功

功能说明：
- 从源读取数据并缓存在p中，并将p中缓存的数据写入目标

示例：
  package main
	
	import (
		"io"
		"os"
		"fmt"
	)
	
	func main() {
		reader, _ := os.Open("src.txt")
		writer, _ := os.Create("dst.txt")
		teeReader := io.TeeReader(reader, writer)
		var n, total int
		var err error
		p := make([]byte, 20)
		for {
			n, err = teeReader.Read(p)
			total = total + n
			if err == nil {
				fmt.Println("Read and write value", string(p[0:n]))
				fmt.Println("Read and write count", total)
			} 
			if err == io.EOF {
				fmt.Println("Read and write end total", total)
				break
			}
		}
	}
