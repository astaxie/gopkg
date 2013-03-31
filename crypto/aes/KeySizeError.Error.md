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
		"crypto/aes"
		"fmt"
	)
	
	func main() {
		// 密钥，这个密钥的长度是错误的，AES要求只能是16，24或32个字节，本例是26个字节
		key := "1234567890abcdef12345678ab"
		// 这个err是KeySizeError类型
		_, err := aes.NewCipher([]byte(key))
		if err != nil {
			fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err.Error())
		}
	}
