# func ParseInt(s string, base int, bitSize int) (i int64, err error)

参数列表

- str     可以表示int64值的字符串
- base    进制 (2 to 36) 
- bitSize 精度 0、8、16、32、64对应int、int8、int16、int32、int64

返回值：

- i       通过str转换的int64值.
- err     当str无法转换为init64值返回错误，否则为nil.

功能说明：

- 尝试按照base指定的进制(2-36)将表示int64值的字符串转换为int64值。如果base为0,则参考s的格式自动确定进制，如："0x"是16进制，"O"是8进制，其余是10进制。
- 返回err是*NumError格式，并且err.Num = s。如果s格式错误，则err.Err = ErrSyntax，并且i将返回0.如果转换的数值超过bitSize指定的带符号数值（signed integer），i将返回符合该bitSize的最大值符号数值。

代码实例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        fmt.Println(strconv.ParseInt("9223372036854775807",10,64))
        fmt.Println(strconv.ParseInt("9223372036854775808",10,64))
        fmt.Println(strconv.ParseInt("0xa",0,64))
    }

代码输出：

    9223372036854775807 <nil>
    9223372036854775807 strconv.ParseInt: parsing "9223372036854775808": value out of range
    10 <nil>