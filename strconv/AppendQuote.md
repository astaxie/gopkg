# func AppendQuote(dst []byte, s string) []byte

参数列表

- dst   原列表
- s     需要append到列表的字符串

返回值:

- []byte  返回列表

功能说明：

- 将浮点数f转换为字符串值，并将转换结果追加到dst的尾部，返回追加后的[]byte。fmt为显示格式。prec是精确的位数，bitSize是64或者32

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
