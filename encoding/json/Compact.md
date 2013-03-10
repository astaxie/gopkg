##func Compact(dst *bytes.Buffer, src []byte) error

参数列表:

- dst 表示字符缓冲区指针
- src 表示JSON格式的字符串切片

返回值:

- 返回error错误信息

功能说明:

这个函数主要是用于将JSON格式的src追加到dst中，正确则返回nil，如果发生错误则返回error信息

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
			"Job":"Programmer"
			}`)
		err := json.Compact(dst, src)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(dst)
	}
