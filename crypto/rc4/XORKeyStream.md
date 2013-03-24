## func (c *Cipher) XORKeyStream(dst, src []byte)

参数列表

- dst 目标
- src 源

返回值：

- 无

功能说明：

XORKeyStream sets dst to the result of XORing src with the key stream. Dst and src may be the same slice but otherwise should not overlap.  
XORKeyStream方法设置dst为src与密钥流的xor(异或)运算结果。dst和src应该是同一的slice但不应重叠。

代码实例：

具体使用参见[func NewCipher(key []byte) (*Cipher, error)](NewCipher.md)
