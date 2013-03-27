# func (w PipeWriter) CloseWithError(err error) error

返回值：关闭是否成功，nil代表成功

功能说明：
- 关闭管道，关闭时正在进行的Read操作将返回参数传入的异常，若管道内仍有未读取的数据，后续仍可正常读取

示例：
  package main;
	
	import (
		"io"
		"fmt"
		"errors"
	)
	
	func main() {
		reader, writer := io.Pipe()
		inputData := []byte("1234567890")
		go writer.Write(inputData)
		var customErr = errors.New("Custom Error For CloseWithError")
		writer.CloseWithError(customErr)
		outputData := make([]byte, 10)
		n, err := reader.Read(outputData)
		fmt.Println("executing read return", err)
		fmt.Println("executing read reads number", n)
		n, _ = reader.Read(outputData)
		fmt.Println(string(outputData))
		fmt.Println("next read number", n)
	}
