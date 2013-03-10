## func (x *Index) FindAllIndex(r *regexp.Regexp, n int) (result [][]int)
参数列表

- r 正则表达式
- n 查找返回结果的个数

返回值

- 返回int二维数组类型

功能说明： 根据正则表达式查找所有的索引，并返回匹配结果在数据中的位置（结果已排序）

代码示例

	package main

	import (
		"fmt"
		"index/suffixarray"
		"regexp"
	)

	func main() {
		data := []byte("aaaaaaa")
		index := suffixarray.New(data)
		str := "[a+]"
		r, _ := regexp.Compile(str)
		res := index.FindAllIndex(r, 1)
		fmt.Println(res) //[[0 1]]
		res = index.FindAllIndex(r, 3)
		fmt.Println(res) //[[0 1][1 2][2 3]]
	}
