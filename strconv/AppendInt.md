# func AppendInt(dst []byte, i int64, base int) []byte

参数列表

- dst   表示原列表
- i     表示需要添加的int64值
- base  表示进制数

返回值：

- 返回[]byte 表示原列表添加数值后新生成的列表

功能说明：

类似AppendFloat，只能追加int类型，base表示int表示的进制数，返回追加后的 []byte

代码实例：

    package main

    import (
        "fmt"
        "strconv"
    )

    func main() {
        newlist := strconv.AppendInt(make([]byte, 0), 123000, 10)
        fmt.Println(newlist) //[49 50 51 48 48 48]
        newlist = strconv.AppendInt(make([]byte, 0), 8, 2)
        fmt.Println(newlist) //[49 48 48 48]
    }
