# func AppendUint(dst []byte, i uint64, base int) []byte

参数列表

- dst   原列表
- i     需要追加到列表的unit64值
- base  unit64的进制

返回值:

- []byte  返回列表

功能说明：

- 类似AppendInt，只能追加uint类型，base表示uint表示的进制数，返回追加后的 []byte

代码示例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        newlist := strconv.AppendUint(make([]byte, 0), 10,2) // 1010
        fmt.Println(newlist)
        newlist = strconv.AppendUint(make([]byte, 0), 10,8) // 12
        fmt.Println(newlist)
    }

代码输出：

    [49 48 49 48]
    [49 50]
