# func ParseBool(str string) (value bool, err error)

参数列表

- str     可以表示bool值的字符串

返回值：

- value   通过str转换的bool值
- err     当str无法转换为bool返回错误，否则为nil

功能说明：

- 尝试将表示bool值的字符串转换为bool值。str可以是1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False，其他的字符将返回错误。

代码实例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        fmt.Println(strconv.ParseBool("1"))
        fmt.Println(strconv.ParseBool("0"))
        fmt.Println(strconv.ParseBool("a"))
    }

代码输出：

    true <nil>
    false <nil>
    false strconv.ParseBool: parsing "a": invalid syntax