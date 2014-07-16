# func AppendQuoteRuneToASCII(dst []byte, r rune) []byte

参数列表

- dst   原列表
- r     需要append到列表的字符

返回值:

- []byte  返回列表

功能说明：

- 将符文s转换为单引号引起来的字符串，非ASCII字符将转换为ASCII，并将结果追加到dst的尾部，返回追加后的[]byte。其中的特殊字符将被转换为转义字符

代码示例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        newlist := strconv.AppendQuoteRune(make([]byte, 0), 'a')
        fmt.Println(newlist)
        newlist = strconv.AppendQuoteRune(make([]byte, 0), '\'')
        fmt.Println(newlist)
        newlist = strconv.AppendQuoteRune(make([]byte, 0), '中')
        fmt.Println(newlist)
        newlist = strconv.AppendQuoteRuneToASCII(make([]byte, 0), 'a')
        fmt.Println(newlist)
        newlist = strconv.AppendQuoteRuneToASCII(make([]byte, 0), '\'')
        fmt.Println(newlist)
        newlist = strconv.AppendQuoteRuneToASCII(make([]byte, 0), '中')
        fmt.Println(newlist)
    }

代码输出：

    [39 97 39]
    [39 92 39 39]
    [39 228 184 173 39]
    [39 97 39]
    [39 92 39 39]
    [39 92 117 52 101 50 100 39]
