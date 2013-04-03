# func FieldsFunc(s string, f func(rune) bool) []string

参数列表

- s 表示需要判断的主串
- f 表示一个函数，该函数参数是rune，返回值是bool，如果rune符合f函数的逻辑那么返回true，否者返回false

返回值：

- 返回[]string 

功能说明：

该函数实现的功能：s的每一个字符传入函数f，如果f返回true，那么按照该字符进行分割（该字符不保留），继续下一个字符，以此类推直到最后，如果返回的都是为false或者s为空，那么将返回空的字符串slice

代码实例：

	package main
	
	import (
		"fmt"
		"strings"
	)
	
	func main() {
		a := strings.FieldsFunc("astaxie", splitfunc)
		fmt.Println(a) //输出：[asta ie]
		//如果把下面的函数t字符改成，那么将返回空slice
	}
	
	func splitfunc(a rune) bool {
		if a > 't' {
			return true
		}
		return false
	}

