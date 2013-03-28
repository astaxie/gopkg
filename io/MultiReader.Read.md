# func (mr *multiReader) Read(p []byte) (n int, err error)

参数：
- p 存放读取结果的比特数组

返回值：
- n 读取字节数
- err 读取是否成功，nil代表成功

功能说明：
- 在MultiReader中读取数据（连续从多个封装的Reader中读取），并存入p中；读完最后一个封装的Reader后返回io.EOF，所以io.EOF代表读取完毕

示例：
  package main;
	
	import (
		"io"
		"fmt"
		"os"
	)
	
	func main() {
		reader1, _ := os.Open("src1.txt")
		reader2, _ := os.Open("src2.txt")
		multiReader := io.MultiReader(reader1, reader2)
		var n, total int
		var err error
		p := make([]byte, 15)
		for {
			n, err = multiReader.Read(p)
			if err == io.EOF {
				fmt.Println("Read end total", total)
				break
			}
			total = total + n
			fmt.Println("Read value:", string(p[0:n]))
			fmt.Println("Read count:", n)
		}
	}
