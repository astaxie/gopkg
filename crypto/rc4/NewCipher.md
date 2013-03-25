## func NewCipher(key []byte) (*Cipher, error)

参数列表

- key 密钥

返回值：

- 返回*Cipher(rc4的Cipher类型), error

功能说明：

NewCipher creates and returns a new Cipher. The key argument should be the RC4 key, at least 1 byte and at most 256 bytes.  
NewCipher方法创建并返回一个新的Cipher类型。参数key应该是RC4的密钥，最少1个字节，最多256个字节。

代码实例：

  	package main
	
	import (
		"crypto/rc4"
		"fmt"
	)
	
	// RC4为对称加密算法
	func main() {
		// 明文
		src := []byte("Hello, world!")
		// 密钥(长度：1到256个字节)
		key := []byte("12345")
	
		cipher, err := rc4.NewCipher(key)
		if err != nil {
			fmt.Println("rc4.NewCipher error:" + err.Error())
		}
	
		encrypted := make([]byte, len(src))
		cipher.XORKeyStream(encrypted, src)
		fmt.Printf("Encrypting %s : %v -> %v\n", src, []byte(src), encrypted)
	
		decrypted := make([]byte, len(encrypted))
		// 下面这个很奇怪? 不知道为什么解密时也得NewCipher一次，上面已经创建了....
		cipher, err = rc4.NewCipher(key)
		if err != nil {
			fmt.Println("rc4.NewCipher error:" + err.Error())
		}
		cipher.XORKeyStream(decrypted, encrypted)
		fmt.Printf("Decrypting %v -> %v : %s\n", encrypted, decrypted, decrypted)
	}
