## func ConstantTimeSelect(v, x, y int) int

参数列表

- v, x, y 整型

返回值：

- int 整型

功能说明：

ConstantTimeSelect returns x if v is 1 and y if v is 0. Its behavior is undefined if v takes any other value.  
如果v等于1， 返回x，如果v等于0，返回y。如果v是其他值，行为未被定义。

代码实例：

  	package main
	
	import (
		"crypto/subtle"
		"fmt"
	)
	
	func main() {
		fmt.Printf("%d\n", subtle.ConstantTimeSelect(1, 2, 3)) // v=1, 返回x, 即2
		fmt.Printf("%d\n", subtle.ConstantTimeSelect(0, 2, 3)) // v=1, 返回y, 即3

		// 不要使用v不是0或1的值，行为未被定义，结果未定。
		fmt.Printf("%d\n", subtle.ConstantTimeSelect(13, 2, 3))
	}
