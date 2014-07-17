# func Quote(s string) string

参数列表

- s     被引用的字符串

返回值：

- 返回string 将s转换为被引用的字符串格式

功能说明：

- 对s两侧添加双引号，并对s中的控制字符和不可打印的字符进行转义。

代码实例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        fmt.Println(strconv.Quote("abc	中文"))
        fmt.Println(strconv.Quote(strconv.Quote("abc	中文")))
    }

代码输出：

    "abc\t中文"
    "\"abc\\t中文\""