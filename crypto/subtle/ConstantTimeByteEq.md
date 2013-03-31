## func ConstantTimeByteEq(x, y uint8) int

参数列表

- x, y 8位无符号整型

返回值：

- int 整型

功能说明：

ConstantTimeByteEq returns 1 if x == y and 0 otherwise.  
如果x与y相等，返回1。否则返回0。

代码实例：

  	package main
	
	import (
		"crypto/subtle"
		"fmt"
	)
	
	func main() {
		fmt.Printf("%d\n", subtle.ConstantTimeByteEq(23, 45)) // x != y, return 0
		fmt.Printf("%d\n", subtle.ConstantTimeByteEq(23, 23)) // x == y, return 1
	}
