# func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)

参数：
- p 存放读取结果的比特数组
- off 偏移量

返回值：
- n 读取的字节数
- err 读取是否成功；nil代表成功

功能说明：
- 从SectionReader中第off+1个字节开始读取数据，并存入p中；若未读到数据返回io.EOF，所以io.EOF代表读取完毕
- 注意：用该方法去取数据时，读取后读取点并不会移动，故多次读取会读取同样结果，无法使用循环读取并用EOF判断结束

示例：
  package main
	
	import (
		"io"
		"os"
		"fmt"
	)
	
	func main() {
		reader, _ := os.Open("src.txt")
		sectionReader := io.NewSectionReader(reader, 5, 10)
		p := make([]byte, 15)
		n, _ := sectionReader.ReadAt(p, 4)
		fmt.Println("Read value:", string(p[0:n]))
		fmt.Println("Read count:", n)
	}
