# func AppendBool(dst []byte, b bool) []byte

参数列表

- dst 表示原列表 
- b   表示需要添加的bool值，true或者false

返回值：

- []byte  返回原列表追加bool后新生成的列表 

功能说明：

- 将布尔值 b 转换为字符串 "true" 或 "false" 然后将结果追加到 dst 的尾部，返回追加后的 []byte

代码实例：

	package main
	
	import (
		"fmt"
		"strconv"
	)
	
	func main() {
		newlist := strconv.AppendBool(make([]byte, 3), false)
		fmt.Println(newlist)//[0 0 0 102 97 108 115 101]
		newlist = strconv.AppendBool(newlist, true)
		fmt.Println(newlist)//[0 0 0 102 97 108 115 101 116 114 117 101]
	}
