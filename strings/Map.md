# func Map(mapping func(rune) rune, s string) string

参数列表

- mapping 处理函数，输入是字符，输出是字符
- s 表示需要处理的主串

返回值：

- 返回string，处理后的字符串

功能说明：

该函数以此读取s中的字符，传入mapping函数，然后返回的字符链接起来，说白了就是字符串的每一个字符通过mapping函数的处理，最后返回处理好的字符串，如果处理不正确，那么就抛弃该字符

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		rot13 := func(r rune) rune {
			switch {
			case r >= 'A' && r <= 'Z':
				return 'A' + (r-'A'+13)%26
			case r >= 'a' && r <= 'z':
				return 'a' + (r-'a'+13)%26
			}
			return r
		}
		fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
		//'Gjnf oevyyvt naq gur fyvgul tbcure...
	}
