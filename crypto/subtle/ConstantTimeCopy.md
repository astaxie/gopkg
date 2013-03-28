## func ConstantTimeCopy(v int, x, y []byte)

参数列表

- v int
- x, y []byte

返回值：

- 无

功能说明：

ConstantTimeCopy copies the contents of y into x if v == 1. If v == 0, x is left unchanged. Its behavior is undefined if v takes any other value.  
如果参数v等于1，将y的内容拷贝到x中。如果参数v等于0，x不变。如果参数v是其他值，行为未被定义。

代码实例：

  	package main
	
	import (
		"crypto/subtle"
		"fmt"
	)
	
	func main() {
	
		y := []byte("lilei")
		x := make([]byte, len(y))
		subtle.ConstantTimeCopy(0, x, y)
		fmt.Printf("%s\n", x) // x = nil*5
	
		subtle.ConstantTimeCopy(1, x, y)
		fmt.Printf("%s\n", x) // x = lilei
		
		x = make([]byte, len(y))
		subtle.ConstantTimeCopy(13, x, y) // 不要如此使用，行为未被定义，结果未知。
		fmt.Printf("%s\n", x)
	
	}
