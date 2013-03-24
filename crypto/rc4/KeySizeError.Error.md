## func (k KeySizeError) Error() string

参数列表

- 无

返回值：

- 返回string

功能说明：

返回密钥长度的错误

代码实例：

  	package main
	
	import (
		"crypto/rand"
		"crypto/rc4"
		"fmt"
	)
	
	func main() {
		// 密钥，这个密钥的长度是错误的，RC4要求最少1个字节，最多256个字节，本例是257个字节
		key := make([]byte, 257)
		rand.Read(key)
		
		// 这个err是KeySizeError类型
		_, err := rc4.NewCipher(key)
		if err != nil {
			fmt.Printf("failed to make RC4 cipher: %s", err)
		}
	}