## func (x *Index) Write(w io.Writer) error
参数列表

- r 字符流

返回值

- 返回错误类型

功能说明： 把字符流写入到索引中

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
		var buf bytes.Buffer //0
		fmt.Println(buf.Len())
		if err := x.Write(&buf); err != nil {
			fmt.Println("failed writing index %s", err)
		}
		fmt.Println(buf.Len()) //24
		
	}