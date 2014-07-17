# func ParseFloat(s string, bitSize int) (f float64, err error)

参数列表

- str     可以表示float64值的字符串
- bitSize 精度 按64位或32位转换为float64值

返回值：

- f       通过str转换的float64值.如果解析出错，返回0
- err     当str无法转换为float64值返回错误，否则为nil

功能说明：

- 尝试按照bitSize指定的精度将表示float64值的字符串转换为float64值。当bitSize指定32位时，返回值仍然是float64值而非float32值,但却可以在不改变数值的情况下将float64转换为float32。
- 如果s是一个float格式或接近float格式的字符串，ParseFloat将返回按照IEEE 754规范舍入的float64值。
- 返回的err是*NumError格式，并且err.Num = s
- 如果s格式不是float格式，返回语法错误 err.Err = ErrSyntax。
- 如果s格式正确但转换成浮点数值后比bitSize指定的精度的最大浮点数大1/2 ULP（unit in the last place），返回值 f = ±Inf, err.Err = ErrRange。
代码实例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        fmt.Println(strconv.ParseFloat("1.0231e2",64))
    }

代码输出：

    102.31 <nil>