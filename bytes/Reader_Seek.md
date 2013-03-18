## func (r *Reader) Seek(offset int64, whence int) (int64, error)

参数列表

- offset 字节偏移
- whence 位置

返回值

- int64 新的偏移字节数
- err 错误

功能说明

- Seek实现了io.Seeker接口，用于设置下次读或写操作的位置，返回新的偏移位置字节数和错误（如果有的话）。
- whence的含义为：
	- 0：从起始位置计算offset
	- 1: 从当前位置计算offset
	- 2：从尾部位置计算offset

代码示例

	package main
	
	import (
		"bytes"
		"fmt"
	)
	
	func main() {
		b := bytes.NewReader([]byte("123456"))
		
	}

代码输出
	
	
