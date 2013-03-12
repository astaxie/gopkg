## func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

返回值

- line 对应于一行的字节切片
- isPrefix 返回的line是否因为太长而被裁减
- err 错误

功能说明

- ReadLine是一个低级的用于读取一行数据的原语。大多数调用者应该使用ReadBytes('\n')或者ReadString('\n')。ReadLine试图返回一行，不包括结尾的回车字符。如果一行太长了（超过缓冲区长度），isPrefix会设置为true并且只返回前面的数据，剩余的数据会在以后的调用中返回。当返回最后一行数据时，isPrefix会设置为false。返回的字节切片只在下一次调用ReadLine前有效。ReadLine或者返回一个非空的字节切片或者返回一个错误，但它们不会同时返回。

代码示例

	package main
	
	import (
		"bufio"
		"bytes"
		"fmt"
	)
	
	func main() {
		rb := bytes.NewBuffer([]byte("123\r\n456"))
		r := bufio.NewReader(rb)
		line, prefix, err := r.ReadLine()
		if err == nil {
			fmt.Printf("%v, %s, %v\n", line, string(line), prefix)
		}
	}
	
代码输出
	
	[49 50 51], 123, false
