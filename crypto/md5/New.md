## func New() hash.Hash

参数列表

- 无

返回值：

- 返回hash.Hash类型

功能说明：

New returns a new hash.Hash computing the MD5 checksum.  
New方法返回一个hash.Hash类型，用于计算MD5的校验和。

代码实例：

  	package main
	
	import (
		"crypto/md5"
		"fmt"
		"io"
	)
	
	// 官网的例子：http://golang.org/pkg/crypto/md5/
	func main() {
		h := md5.New()
		io.WriteString(h, "The fog is getting thicker!")
		io.WriteString(h, "And Leon's getting laaarger!")
		fmt.Printf("%x", h.Sum(nil))
	}
	// e2c569be17396eca2a2e3c11578123ed

