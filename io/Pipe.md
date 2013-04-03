# func Pipe() (*PipeReader, *PipeWriter)

返回值：
- PipeReader：管道读取器
- PipeWriter：管道写入器

功能说明：
- 创建一个管道，并返回它的读取器和写入器

示例：
  package main
	
	import (
		"io"
		"fmt"
		"reflect"
	)
	
	func main() {
		reader, writer := io.Pipe()
		fmt.Println(reflect.TypeOf(reader))
		fmt.Println(reflect.TypeOf(writer))
	}
	
	
	
