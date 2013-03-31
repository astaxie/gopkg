# func (s *SectionReader) Read(p []byte) (n int, err error)

参数：
- p 存放读取结果的比特数组

返回值：
- n 读取的字节数
- err 读取是否成功；nil代表成功

功能说明：
- 在SectionReader中读取数据，并存入p中；若未读到数据返回io.EOF，所以io.EOF代表读取完毕

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
		var n, total int
		var err error
		p := make([]byte, 15)
		for {
			n, err = sectionReader.Read(p)
			if err == io.EOF {
				fmt.Println("Find EOF so end total", total)
				break
			}
			total = total + n
			fmt.Println("Read value:", string(p[0:n]))
			fmt.Println("Read count:", n)
		}
	}
