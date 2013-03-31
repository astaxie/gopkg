## func (x *Index) Lookup(s []byte, n int) (result []int)

参数列表

- s 需要查找的byte类型数组
- n 返回结果的个数，n<0 则返回全部

返回值

- 返回int数组类型

功能说明：根据输入byte数组查找索引，并返回匹配结果在数据中的位置（结果未排序）

代码示例

	package main

	import (
		"fmt"
		"index/suffixarray"
	)
	
	func main() {
		data := []byte("aaaaaaa")
		index := suffixarray.New(data)
		str := []byte("a")
		res := index.Lookup(str, 1)
		fmt.Println(res) //[6]
		res = index.Lookup(str, 3)
		fmt.Println(res) //[6 5 4]
		res = index.Lookup(str, -1)
		fmt.Println(res) //[6 5 4 3 2 1 0]
	}
