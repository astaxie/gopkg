## func NewReader(b []byte) *Reader

参数列表

- b 字节切片

返回值

- *Reader

功能说明

- NewReader创建一个Reader，数据为b。

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewReader([]byte("12345"))
		fmt.Println(b.Len())
	}

代码输出
	
	5
