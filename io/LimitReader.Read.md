# func (l *LimitedReader) Read(p []byte) (n int, err error)

参数：
- p 存放读取结果的比特数组

返回值：
- n 读取的字节数
- err 读取是否成功；nil代表成功

功能说明：
- 在LimitedReader中读取数据，并存入p中；每次读取后，限定字节数=限定字节数-len(p)；若len(p)>限定字节数，len(p):=限定字节数；若限定字节数<=0,返回io.EOF，所以io.EOF代表读取完毕

示例：
  package main;
	
	import (
		"io"
		"fmt"
		"os"
	)
	
	func main() {
		reader, _ := os.Open("readFile.txt")
		limitReader := io.LimitReader(reader, 20)
		var n, total int
		var err error
		p := make([]byte, 15)
		for {
			n, err = limitReader.Read(p)
			if err == io.EOF {
				fmt.Println("Read end total", total)
				break
			}
			total = total + n
			fmt.Println("Read value:", string(p[0:n]))
			fmt.Println("Read count:", n)
		}
	}
