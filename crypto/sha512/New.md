## func New() hash.Hash

参数列表

- 无

返回值：

- 返回hash.Hash

功能说明：

New returns a new hash.Hash computing the SHA512 checksum.  
返回hash.Hash，可以进行SHA512（安全哈希算法）计算校验和值。

代码实例：

  	package main
	
	import (
		"crypto/sha512"
		"fmt"
		"io"
	)
	
	func main() {
		h := sha512.New()
		io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
		fmt.Printf("% x", h.Sum(nil))
	}
	
	//结果，长度为64个字节
	//9a 98 dd 9b b6 7d 0d a7 bf 83 da 53 13 df f4 fd 60 a4 ba c0 09 4f 1b 05 63 36 90 ff a7 f6 d6 1d e9 a1 d4 f8 61 79 37 d5 60 83 3a 9a aa 9c ca fe 3f d2 4d b4 18 d0 e7 28 83 35 45 ca dd 3a d9 2d
