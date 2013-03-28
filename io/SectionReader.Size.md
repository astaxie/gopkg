# func (s *SectionReader) Size() int64

返回值：可读取字节数量

功能说明：
- 获得SectionReader可读取字节数量，不受Read影响

示例：
  package main
	
	import (
		"io"
		"os"
		"fmt"
	)
	
	func main() {
		reader, _ := os.Open("src.txt")
		sectionReader := io.NewSectionReader(reader, 5, 20)
		fmt.Println("Can read count:", sectionReader.Size())
		p := make([]byte, 10)
		n, _ := sectionReader.Read(p)
		fmt.Println("Read count:", n)
		fmt.Println("Can read count:", sectionReader.Size())
	}
