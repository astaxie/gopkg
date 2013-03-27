# func MultiWriter(writers ...Writer) Writer

参数：
- writers 不定参数，为若干写入器

返回值：MultiWriter

MultiWriter说明：
- 封装多个Writer，统一的Write操作可以将数据同时写入多目标

功能说明：
- 获得一个可以对多Writer同时写入的Writer

示例：
  package main
	
	import (
		"io"
		"fmt"
		"os"
		"reflect"
	)
	
	func main() {
		writer1, _ := os.Create("dst1.txt")
		writer2, _ := os.Create("dst2.txt")
		multiWriter := io.MultiWriter(writer1, writer2)
		fmt.Println(reflect.TypeOf(multiWriter))
	}
