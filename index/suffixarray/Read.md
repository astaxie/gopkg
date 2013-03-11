## func (x *Index) Read(r io.Reader) error
参数列表

- r 字符流

返回值

- 返回错误类型

功能说明： 从一个字符流中读取数据到索引

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
		"index/suffixarray"
	)
	
	func main() {
		data := []byte("aa")
		x := suffixarray.New(data)
		var buf bytes.Buffer
		if err := x.Write(&buf); err != nil {
			fmt.Println("failed writing index %s", err)
		}
		size := buf.Len()
		var y suffixarray.Index
		if err := y.Read(&buf); err != nil {
			fmt.Println("failed reading index %s ", err)
		}
		fmt.Println(y.Bytes()) //[97 97]
	}