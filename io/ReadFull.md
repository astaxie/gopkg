# func ReadFull(r Reader, buf []byte) (n int, err error)

参数：
- r 源读取器
- buf 存放读取结果的比特数组

返回值：
- n 读取的字节数
- err 读取是否成功；nil代表成功

功能说明：
- 在r中读取len(buf)字节数据，并存放在buf中

可能的异常：
- io.ErrUnexpectedEOF：读取的数据不足len(buf)字节
- io.EOF：读取完毕

示例：
  package main
	
	import (
		"io"
		"os"
		"fmt"
	)
	
	func main() {
		reader, _ := os.Open("src.txt")
		var n, total int
		var err error
		p := make([]byte, 20)
		fmt.Println("Full count is:", len(p))
		for {
			n, err = io.ReadFull(reader, p)
			if err == nil {
				fmt.Println("Read enough value:", string(p))
			}
			if err == io.ErrUnexpectedEOF {
				fmt.Println("Read fewer value:", string(p[0:n]))
			}
			if err == io.EOF {
				fmt.Println("Read end total", total)
				break
			}
			total = total + n
		}
	}
