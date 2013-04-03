# func (s *SectionReader) Seek(offset int64, whence int) (ret int64, err error)

参数：
- offset 偏移量
- whence 设定选项 0:读取起始点，1:当前读取点，2:结束点(不好用)，其他：将抛出Seek: invalid whence异常

返回值：
- ret 当前读取点相对读取起始点的偏移量
- err 便宜是否成功；nil代表成功

功能说明：
- 对SectionReader的读取起始点、当前读取点、结束点进行偏移

示例：
  package main
	
	import (
		"io"
		"os"
		"fmt"
	)
	
	func main() {
		reader, _ := os.Open("src.txt")
		sectionReader := io.NewSectionReader(reader, 2, 20)
		sectionReader.Seek(2, 0)
		p := make([]byte, 10)
		n, _ := sectionReader.Read(p)
		fmt.Println("First read value:", string(p[0:n]))
		fmt.Println("First read count:", n)
		a, _ := sectionReader.Seek(2, 1)
		fmt.Println("off - base is", a)
		n, _ = sectionReader.Read(p)
		fmt.Println("Second Read value:", string(p[0:n]))
		fmt.Println("Second Read count:", n)
		sectionReader.Seek(8, 2)
		n, _ = sectionReader.Read(p)
		fmt.Println("Third read value:", string(p[0:n]))
		fmt.Println("Third read count:", n)
	}
