## func (enc *Encoding) Encode(dst, src []byte)

参数列表：

- dst 编码缓冲区
- src 要编码的字符串切片

返回值：

无

功能说明：

对输入的字符串进行 base64 编码

代码实例：

    package main

    import (
        "fmt"
        "encoding/base64"
    )

    func Example1() {

        fmt.Println("---=== Example1 ===---")

        // 要编码的字符串
        src := "this is a test string."
        // 计算编码缓冲区需要的长度
        l := base64.StdEncoding.EncodedLen(len(src))

        // 创建足够长度的缓冲区
        dst := make([]byte, l)

        // 编码
        base64.StdEncoding.Encode(dst, []byte(src))

        fmt.Println(string(dst) == "dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg==") // true
        fmt.Println(l == len(dst)) // true
        fmt.Println(l%4 == 0) // true

    }

    func Example2() {

        fmt.Println("---=== Example2 ===---")

        // 编码缓冲区需要的长度会被填充到4的倍数
        for ii := 0; ii < 20; ii++ {
            l := base64.StdEncoding.EncodedLen(ii)
            fmt.Println(l, l%4 == 0) // true
        }
    }

    func main() {

        Example1()
        Example2()

    }
