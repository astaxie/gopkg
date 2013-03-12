## func New() hash.Hash

参数列表

- 无

返回值：

- 返回hash.Hash

功能说明：

New returns a new hash.Hash computing the SHA256 checksum.  
返回hash.Hash，可以进行SHA256（安全哈希算法）计算校验和值。

代码实例：

  	package main
	
	import (
		"crypto/sha256"
		"fmt"
		"io"
	)
	
	func main() {
		h := sha256.New()
		io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
		fmt.Printf("% x", h.Sum(nil))
	}

	//结果，长度为32个字节
	//80 01 f1 90 df b5 27 26 1c 4c fc ab 70 c9 8e 80 97 a7 a1 92 21 29 bc 40 96 95 0e 57 c7 99 9a 5a
