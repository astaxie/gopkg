## func New() hash.Hash

参数列表

- 无

返回值：

- 返回hash.Hash类型

功能说明：

New returns a new hash.Hash computing the SHA1 checksum.  
New方法返回一个hash.Hash类型，用于计算SHA1的校验和。

代码实例：

  	package main
	
	import (
		"crypto/sha1"
		"fmt"
		"io"
	)
	// 官网的例子：http://golang.org/pkg/crypto/sha1/
	func main() {
		h := sha1.New()
		io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
		fmt.Printf("% x", h.Sum(nil))
	}

	// 59 7f 6a 54 00 10 f9 4c 15 d7 18 06 a9 9a 2c 87 10 e7 47 bd

