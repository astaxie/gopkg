# func WriteString(w Writer, s string) (n int, err error)

参数：
- w 目标写入器
- s 写入的字符串

返回值：
- n 写入字节数
- err 写入是否成功，nil代表成功

功能说明：
- 将字符串s写入到w中

示例：
  package main
	
	import (
		"io"
		"fmt"
		"os"
	)
	
	func main() {
		writer, _ := os.Create("dst.txt")
		data := "write data.."
		n, err := io.WriteString(writer, data)
		if err == nil {
			fmt.Println("Success write", n, "byte")
		}
	}
