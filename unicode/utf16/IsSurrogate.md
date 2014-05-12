# func IsSurrogate(r rune) bool 

参加列表

- r 

返回值:

- 成功返回true
- 失败返回false

功能说明:

>IsSurrogate 判断 r 是否为代理区字符
>
>两个代理区字符可以用来组合成一个 utf16 编码
>
>代理区字符分别有三个 0xd800, 0xdc00, 0xdfff 
> 如果r是代理区字符，则返回true
>
> 如果r不是代理区字符，则返回false
>

代码实例:

	package main
	
	import ( 
        "fmt"
        "unicode/utf16"
        )

        func main() {
        fmt.Println(utf16.IsSurrogate(0xD800)) // true
	fmt.Println(utf16.IsSurrogate(0xDC00)) // true
	fmt.Println(utf16.IsSurrogate(0xDFFF)) // true
        }
        

