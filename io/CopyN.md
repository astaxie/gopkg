# func CopyN(dst Writer, src Reader, n int64) (written int64, err error)

参数：
- dst 拷贝目标写入器
- src 拷贝源读取器
- n 拷贝的字节数

返回值：
- written 拷贝字节数
- err 拷贝是否成功；nil代表拷贝成功

功能说明：
- 向dst拷贝src的n比特数据

可能的异常：
- io.EOF：当n>src字节数时，拷贝src全部数据并返回src字节数
- io.ErrShortWrite:写入数据不等于读取数据

示例：
  package main
	
	import (
		"io"
		"fmt"
		"os"
	)
	
	func main() {
		srcFile, _ := os.Open("copySrc.txt")
		destFile, _ := os.Create("copyDest.txt")
		written, err := io.CopyN(destFile, srcFile, 15)
		if err == nil {
			fmt.Println("Copy Success! total", written, "bytes")
		}
		if err == io.EOF {
			fmt.Println("Copy all total", written, "bytes")
		}
	}
	
	
