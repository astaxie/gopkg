## func (r *Reader) Seek(offset int64, whence int) (int64, error)

参数列表

- offset 相对whence的字节偏移（从0开始计算）
- whence 位置

返回值

- int64 新的偏移字节数（从0开始计算）
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
	
	var (
		data = []byte("123456")
	)
	
	func seekhead() {
		fmt.Println("seekhead:")
		b := bytes.NewReader(data)
		// 把位置移到起始位置+1字节处，即0+1==1，对应数据为'2'
		fmt.Println(b.Seek(1, 0))
		c, _ := b.ReadByte()
		fmt.Println(string(c))
	}
	
	func seekcur() {
		fmt.Println("seekcur:")
		b := bytes.NewReader(data)
		// 连续读取两个字节后，当前偏移位置为2字节
		b.ReadByte()
		b.ReadByte()
		// 把位置移到当前位置+1字节处，即2+1==3，对应数据为'4'
		fmt.Println(b.Seek(1, 1))
		c, _ := b.ReadByte()
		fmt.Println(string(c))
	}
	
	func seektail() {
		fmt.Println("seektail:")
		b := bytes.NewReader(data)
		// 把位置移到尾部位置-2字节处，即6-2==4，对应数据为'5'
		fmt.Println(b.Seek(-2, 2))
		c, _ := b.ReadByte()
		fmt.Println(string(c))
	}
	
	func main() {
		seekhead()
		seekcur()
		seektail()
	}

代码输出
	
	seekhead:
	1 <nil>
	2
	seekcur:
	3 <nil>
	4
	seektail:
	4 <nil>
	5

