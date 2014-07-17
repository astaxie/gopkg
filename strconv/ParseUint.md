# func ParseUint(s string, base int, bitSize int) (n uint64, err error)

参数列表

- str     可以表示int64值的字符串
- base    进制 (2 to 36) 
- bitSize 精度 0、8、16、32、64对应uint、uint8、uint16、uint32、uint64

返回值：

- i       通过str转换的uint64值.
- err     当str无法转换为uinit64值返回错误，否则为nil.

功能说明：

- ParseUint is like [ParseInt](ParseInt.md) but for unsigned numbers.

代码实例：

    package main
    
    import (
        "fmt"
        "strconv"
    )
    
    func main() {
        fmt.Println(strconv.ParseUint("18446744073709551615",10,64))
        fmt.Println(strconv.ParseUint("18446744073709551616",10,64))
        fmt.Println(strconv.ParseUint("0xa",0,64))
    }

代码输出：

    18446744073709551615 <nil>
    18446744073709551615 strconv.ParseUint: parsing "18446744073709551616": value out of range
    10 <nil>