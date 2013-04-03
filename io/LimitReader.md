# func LimitReader(r Reader, n int64) Reader

参数：
- r 读取器
- n 限定字节数

返回值：LimitReader对象

LimitReader对象说明：
- 是对Reader对象的一种封装，限定Reader数据的读取数量

功能说明：
- 获得一个只能从r读取n比特数据的Reader

示例：
  package main
	
	import (
		"io"
		"fmt"
		"os"
		"reflect"
	)
	
	func main() {
		reader, _ := os.Open("readFile.txt")
		limitReader := io.LimitReader(reader, 20)
		fmt.Println(reflect.TypeOf(limitReader))
	}
