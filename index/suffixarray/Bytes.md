## func (x *Index) Bytes() []byte
参数列表

- 无

返回值

- 返回[]byte类型

功能说明： 获取Index的数据

代码示例

	package main
	
	import (
		"fmt"
		"index/suffixarray"
	)
	
	func main() {
		data := []byte("aaaa")
		index := suffixarray.New(data) 
		fmt.Println(index.Bytes()) //[97,97,97,97]
	}