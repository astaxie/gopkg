## func ConstantTimeEq(x, y int32) int

参数列表

- x, y 32位整型

返回值：

- int 整型

功能说明：

ConstantTimeEq returns 1 if x == y and 0 otherwise.  
如果x与y相等，返回1。否则返回0。

代码实例：

  	package main
	
	import (
		"crypto/subtle"
		"fmt"
	)
	
	func main() {
		fmt.Printf("%d\n", subtle.ConstantTimeEq(0, 1)) // 0
		fmt.Printf("%d\n", subtle.ConstantTimeEq(-1, -1)) // 1
		fmt.Printf("%d\n", subtle.ConstantTimeEq(-1, 1)) // 0
	}
