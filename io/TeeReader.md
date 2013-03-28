# func TeeReader(r Reader, w Writer) Reader

参数：
- r 源读取器
- w 目标写入去

返回值：TeeReader对象

TeeReader对象说明：
- 是对Reader对象的一种封装，从r中读取写入并写入w中

功能说明：
- 创建一个TeeReader对象

示例：
  package main
	
	import (
		"io"
		"os"
		"fmt"
		"reflect"
	)
	
	func main() {
		reader, _ := os.Open("copySrc.txt")
		writer, _ := os.Create("copyDest.txt")
		teeReader := io.TeeReader(reader, writer)
		fmt.Println(reflect.TypeOf(teeReader))
	}
