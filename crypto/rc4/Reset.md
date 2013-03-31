## func (c *Cipher) Reset()

参数列表

- 无

返回值：

- 无

功能说明：

Reset zeros the key data so that it will no longer appear in the process's memory.  
Reset（重置）方法清空密钥数据，所以它将不会在进程的内存中出现。

代码实例：

  	package main
	
	import (
		"crypto/rc4"
		"fmt"
	)
	
	func main() {
		// 明文
		src := []byte("Hello, world!")
		// 密钥(长度：1到256个字节)
		key := []byte("12345")
	
		cipher, _ := rc4.NewCipher(key)
	
		encrypted := make([]byte, len(src))
		cipher.XORKeyStream(encrypted, src)
		fmt.Printf("Encrypting %s : %v -> %v\n", src, []byte(src), encrypted)
		
		// 重置，再进行异或运算已经无效
		cipher.Reset()
		decrypted := make([]byte, len(encrypted))
		cipher.XORKeyStream(decrypted, encrypted)
		fmt.Printf("Decrypting %v -> %v : %s\n", encrypted, decrypted, decrypted)
	}
