## func NewTripleDESCipher(key []byte) (cipher.Block, error)

参数列表

- key 密钥

返回值：

- 返回cipher.Block, error

功能说明：

NewTripleDESCipher creates and returns a new cipher.Block.  
创建并返回cipher.Block。  
3DES是对每个数据块进行3次DES加密算法，密钥更长而已。

代码实例：

	package main
	
	import (
		"bytes"
		"crypto/cipher"
		"crypto/des"
		"fmt"
		"os"
	)
	
	// 本例使用3DES演示使用CBC模式（CBC模式必须有填充）
	// 算法/模式/填充
	// 3DES/CBC/PKCS5Padding
	func main() {
		// 消息明文
		src := []byte("hello, world")
		// 密钥，长度为24字节
		key := "12345678" + "12345678" + "12345678"
		// 初始向量，长度大于8个字节
		var iv = []byte("12345678")
		// 得到块，用于加密和解密
		block, err := des.NewTripleDESCipher([]byte(key))
		if err != nil {
			fmt.Printf("Error: NewTripleDESCipher(%d bytes) = %s", len(key), err)
			os.Exit(1)
		}
		fmt.Printf("des.NewTripleDESCipher(%d bytes)\n", len(key))
	
		// 加密，使用CBC模式(密码块链接模式)
		encrypter := cipher.NewCBCEncrypter(block, iv)
		// CBC要求明文必须是16字节的倍数，否则需要Padding
		src = PKCS5Padding(src, block.BlockSize())
		encrypted := make([]byte, len(src))
		encrypter.CryptBlocks(encrypted, src)
		fmt.Printf("Encrypting %s : %v -> %v\n", src, []byte(src), encrypted)
	
		// 解密
		decrypter := cipher.NewCBCDecrypter(block, iv)
	
		decrypted := make([]byte, len(src))
		decrypter.CryptBlocks(decrypted, encrypted)
		decrypted = PKCS5UnPadding(decrypted)
	
		fmt.Printf("Decrypting %v -> %v : %s\n", encrypted, decrypted, decrypted)
	}
	
	// PKCS5Padding
	func PKCS5Padding(src []byte, blockSize int) []byte {
		padding := blockSize - len(src)%blockSize
		padtext := bytes.Repeat([]byte{byte(padding)}, padding)
		return append(src, padtext...)
	}
	
	// PKCS5UnPadding
	func PKCS5UnPadding(src []byte) []byte {
		length := len(src)
		// 去掉最后一个字节 unpadding 次
		unpadding := int(src[length-1])
		return src[:(length - unpadding)]
	}