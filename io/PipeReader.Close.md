# func (r *PipeReader) Close() error

返回值：关闭是否成功，nil代表成功

功能说明：
- 关闭管道，关闭后管道将不能写入

可能的异常：
- 管道关闭后，正在进行或后续的写入Write操作返回ErrClosedPipe

示例：
  package main;
	
	import (
		"io"
		"fmt"
	)
	
	func main() {
		reader, writer := io.Pipe()
		inputData := []byte("1234567890")
		reader.Close()
		c := make(chan error)
		go func() {
			_, err := writer.Write(inputData)
			c <- err
		}()
		if io.ErrClosedPipe == <-c {
			fmt.Println("return ErrClosedPipe")
		}
	}
