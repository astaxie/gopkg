# func AppendQuote(dst []byte, s string) []byte

参数列表

- dst   原列表
- s     需要append到列表的字符串

返回值:

- []byte  返回列表

功能说明：

- 将字符串s转换为双引号引起来的字符串，并将结果追加到dst的尾部，返回追加后的[]byte。其中的特殊字符将被转换为转义字符

代码示例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        newlist := strconv.AppendQuote(make([]byte, 0), "")
        fmt.Println(newlist)
        newlist = strconv.AppendQuote(make([]byte, 0), "abc")
        fmt.Println(newlist)
        newlist = strconv.AppendQuote(make([]byte, 0), "中文")
        fmt.Println(newlist)
        newlist = strconv.AppendQuote(make([]byte, 0), "	") // \t
        fmt.Println(newlist)
    }



代码输出：

    [34 34]
    [34 97 98 99 34]
    [34 228 184 173 230 150 135 34]
    [34 92 116 34]
