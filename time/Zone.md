# func (t Time) Zone() (name string, offset int)

参数列表：

- 无

返回值：

- name 时区简写
- offset 相对与UTC向东的秒偏移

功能说明：

	package main
	
	import (
	    "fmt"
	    "time"
	)
	
	func main() {
	    fmt.Println(time.Now().Zone())
	}
