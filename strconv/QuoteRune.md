# func QuoteRune(r rune) string

参数列表

- r     被引用的rune字符

返回值：

- 返回string 将r转换为被引用的字符格式

功能说明：

- 对s两侧添加单引号，并对r中的控制字符和不可打印的字符进行转义。

代码实例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        fmt.Println(strconv.QuoteRune('	'))
        fmt.Println(strconv.QuoteRune(100))
    }

代码输出：

    '\t'
    'd'