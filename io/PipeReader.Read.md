# func (r *PipeReader) Read(data []byte) (n int, err error)

参数：
- data 存放读取结果的比特数组

返回值：
- n 读取比特数
- err 读取是否成功，nil代表成功

功能说明：
- 从管道中读取数据，并存放到data中

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
