# func EncodeRune( r rune) (r1, r2 rune) 

参数列表:

- r 要编码的字符
- 如果 r < 0x10000, 则无需编码，其utf-16序列就是其自身

返回值列表:
r1：编码后的 UTF-16 代理对的高位码元
r2：编码后的 UTF-16 代理对的低位码元
如果 r 不是有效的 Unicode 字符，或者是代理区字符，或者无需编码
则返回 U+FFFD, U+FFFD

功能说明：

>将字符 r 编码成 UTF-16 代理对

代码实例:

package main

	import (
	"fmt"
	"unicode/utf16"
	)
	func main() {
	r1, r2 := utf16.EncodeRune('\U0010F000')
	fmt.Printf("%x, %x", r1, r2)
	// dbfc, dc00
	}
      
 
