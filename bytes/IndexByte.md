## func IndexByte(s []byte, c byte) int

参数列表

- s 字节切片
- c 字节

返回值

- int s中的位置索引（从0开始）

功能说明

- IndexByte检查c在s中第一次出现的位置索引；如果s中不包含c则返回-1

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		fmt.Println(bytes.IndexByte([]byte("1234"), '3'))
	}

代码输出

	2
