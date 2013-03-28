## func ConstantTimeCompare(x, y []byte) int

参数列表

- x, y []byte

返回值：

- int 整型

功能说明：

ConstantTimeCompare returns 1 if the two equal length slices, x and y, have equal contents. The time taken is a function of the length of the slices and is independent of the contents.  
如果两个长度相等的slices（x与y）并且内容也相等，返回1。所花费的时间与slice的长度有关于内容无关。

注意： 文档没有说两个slices不相等时的返回值情况，注意下面的例子，小心使用

代码实例：

  	package main
	
	import (
		"crypto/subtle"
		"fmt"
	)
	
	func main() {
		fmt.Printf("%d\n", subtle.ConstantTimeCompare([]byte("lilei"), []byte("lilei"))) // 1
		fmt.Printf("%d\n", subtle.ConstantTimeCompare([]byte("lilei"), []byte("lilei2"))) // 1
		fmt.Printf("%d\n", subtle.ConstantTimeCompare([]byte("lilei"), []byte(" lilei"))) // 0
		fmt.Printf("%d\n", subtle.ConstantTimeCompare([]byte("lilei "), []byte("liming"))) // 0
	}
