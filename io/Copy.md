# func Copy(dst Writer, src Reader) (written int64, err error)

参数：
- dst 拷贝目标写入器
- src 拷贝源读取器

返回值：
- written 拷贝字节数
- err 拷贝是否成功；nil代表拷贝成功

功能说明：
向dst拷贝src的全部数据；读取src中数据直到EOF，故不会返回io.EOF

可能的异常：
io.ErrShortWrite:写入数据不等于读取数据

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
		written, err := io.Copy(destFile, srcFile)
		if err == nil {
			fmt.Println("Copy Success! total", written, "bytes")
		}
	}
