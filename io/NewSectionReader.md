# func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader

参数：
- r 源读取器
- off 偏移量（从r的第off+1个字节开始读）
- n 可读取字节数量

返回值：SectionReader对象

功能说明：
- 创建一个SectionReader

SectionReader说明：
- 实现对Reader的部分读取

示例：
  package main
	
	import (
		"io"
		"os"
		"fmt"
		"reflect"
	)
	
	func main() {
		reader, _ := os.Open("src.txt")
		sectionReader := io.NewSectionReader(reader, 5, 10)
		fmt.Println(reflect.TypeOf(sectionReader))
	}
