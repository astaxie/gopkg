##func HTMLEscape(dst *bytes.Buffer, src []byte) error

参数列表:

- dst 表示字符缓冲区指针
- src 表示JSON格式的字符串切片

返回值:

- 无

功能说明:

这个函数主要是用于将JSON格式的src追加到dst中，与Compact不同的是会将<， >， 和 & 编码成 \u003c, \u003e, \u0026，以便于可以安全的嵌入到html的\<script\>标签中,(这么做主要是介于历史原因，浏览器在\<script\>标签中不支持标准的html转义)

代码实例:

	package main

	import (
		"bytes"
		"encoding/json"
		"fmt"
	)
	
	func main() {
		dst := new(bytes.Buffer)
		src := []byte(`{
			"Name":"tony.shao",
			"Age":25,
			"Job":"Programmer<Escaping>"
			}`)
		json.HTMLEscape(dst, src)
		fmt.Println(dst)
	}
