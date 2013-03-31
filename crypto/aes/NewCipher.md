## func NewCipher(key []byte) (cipher.Block, error)

参数列表

- key 密钥

返回值：

- 返回cipher.Block, error

功能说明：

NewCipher creates and returns a new cipher.Block. The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.  
返回cipher.Block，这个块是加密和解密的参数(AES为对称加密算法，即加密和解密需要相同的密钥)。参数key应该是AES的密钥，可以为16, 24, 32字节，分别对应AES-128, AES-192或AES-256。

代码实例：

  	package main
	
	import (
		"crypto/aes"
		"crypto/cipher"
		"fmt"
		"os"
	)
	
	func main() {
		// 消息明文
		src := []byte("hello, world")
		// 密钥，长度可以为16、24、32字节
		key := "1234567890abcdef"
		// 初始向量，长度必须为16个字节(128bit)
		var iv = []byte("abcdef1234567890")[:aes.BlockSize]
		// 得到块，用于加密和解密
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
			os.Exit(1)
		}
		fmt.Printf("NewClipher(%d bytes)\n", len(key))

		// 加密，使用CFB模式(密文反馈模式)，其他模式参见crypto/cipher
		encrypter := cipher.NewCFBEncrypter(block, iv)

		encrypted := make([]byte, len(src))
		encrypter.XORKeyStream(encrypted, src)
		fmt.Printf("Encrypting %s : %v -> %v\n", src, []byte(src), encrypted)

		// 解密
		decrypter := cipher.NewCFBDecrypter(block, iv)

		decrypted := make([]byte, len(src))
		decrypter.XORKeyStream(decrypted, encrypted)
		fmt.Printf("Decrypting %v -> %v : %s\n", encrypted, decrypted, decrypted)
	}
