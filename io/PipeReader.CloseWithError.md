# func (r *PipeReader) CloseWithError(err error) error

参数：
- err 任何实现err接口的对象

返回值：关闭是否成功，nil代表成功

功能说明：
- 关闭管道，关闭后管道将不能写入；管道关闭后，正在进行或后续的写入Write操作返回参数传入的异常

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
		var customErr = errors.New("Custom Error For CloseWithError")
		reader.CloseWithError(customErr)
		c := make(chan error)
		go func() {
			_, err := writer.Write(inputData)
			c <- err
		}()
		fmt.Println(<-c)
	}
