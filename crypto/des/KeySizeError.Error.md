## func (k KeySizeError) Error() string

参数列表

- 无

返回值：

- 返回string

功能说明：

返回密钥长度的错误.

代码实例：

  	package main
	
	import (
		"crypto/des"
		"fmt"
	)
	
	func main() {
		// 密钥，这个密钥的长度是错误的，DES要求只能是8个字节，本例是16个字节
		key := "1234567890abcdef"
		// 这个err是KeySizeError类型
		_, err := des.NewCipher([]byte(key))
		if err != nil {
			fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err.Error())
		}
	}