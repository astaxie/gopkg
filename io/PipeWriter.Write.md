# func (w *PipeWriter) Write(data []byte) (n int, err error)

参数：
- data 写入管道的比特数据

返回值：
- n 写入比特数
- err 写入是否成功，nil代表成功

功能说明：
- 将data比特数组的数据写入管道

示例：
  package main;
	
	import (
		"io"
		"fmt"
	)
	
	func main() {
		reader, writer := io.Pipe()
		inputData := []byte("1234567890")
		go writer.Write(inputData)
		outputData := make([]byte, 10)
		n, _ := reader.Read(outputData)
		fmt.Println(string(outputData))
		fmt.Println("read number", n)
	}
