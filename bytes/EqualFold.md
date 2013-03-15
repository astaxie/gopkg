## func EqualFold(s, t []byte) bool

参数列表

- s, t 要比较的字节切片

返回值

- bool 是否相等

功能说明

- EqualFold把s和t解释成UTF-8字符串，并进行大小写不敏感的比较

代码示例

	package main

	import (
		"bytes"
		"fmt"
	)

	func main() {
		fmt.Println(bytes.EqualFold([]byte("abc"), []byte("abc")))
		fmt.Println(bytes.EqualFold([]byte("abc"), []byte("abd")))
		fmt.Println(bytes.EqualFold([]byte("abc"), []byte("aBc")))
	}


代码输出

	true
	false
	true
