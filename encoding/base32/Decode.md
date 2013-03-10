## func (enc *Encoding) Decode(dst, src []byte) (n int, err error)

参数列表：

- dst 解码缓冲区
- src 要解码的字符串切片

返回值：

- n 解码到缓冲区的字节数
- err 可能的错误

功能说明：

对输入的字符串进行 base32 解码

代码实例：

    package main

    import (
        "fmt"
        "encoding/base32"
    )

    func Example1() {

        fmt.Println("---=== Example1 ===---")

        // 原始字符串
        origin := "this is a test string."

        // 要解码的字符串
        src := "ORUGS4ZANFZSAYJAORSXG5BAON2HE2LOM4XA===="
        // 计算解码缓冲区需要的长度
        maxlen := base32.StdEncoding.DecodedLen(len(src))

        // 创建足够长度的缓冲区
        dst := make([]byte, maxlen)

        // 解码
        n, err := base32.StdEncoding.Decode(dst, []byte(src))
        fmt.Println(err == nil) // true

        // n 小于等于缓冲区长度
        fmt.Println(n <= maxlen) // true

        // n 等于原始字符串的长度
        fmt.Println(n == len(origin)) // true

        // 原始字符串和解码后字符串相同
        fmt.Println(string(dst[0:n]) == origin) // true

    }

    // 测试待解码内容中含有 "\r" "\n" 的情况(会被忽略)
    func Example2() {

        fmt.Println("---=== Example2 ===---")
        origin := "this is a test string."
        src := "ORUGS4ZANFZSAY\r\nJAOR\nSXG5BAON\r2HE2LOM4XA===="
        maxlen := base32.StdEncoding.DecodedLen(len(src))
        dst := make([]byte, maxlen)
        n, err := base32.StdEncoding.Decode(dst, []byte(src))
        fmt.Println(err == nil) // true
        fmt.Println(n <= maxlen) // true
        fmt.Println(n == len(origin)) // true
        fmt.Println(string(dst[0:n]) == origin) // true

    }

    // 测试解码失败的情况
    func Example3() {

        fmt.Println("---=== Example3 ===---")
        origin := "this is a test string."
        src := "ORUGS4ZANFZSA@YJAORSXG5BAON2HE2LOM4XA===="
        maxlen := base32.StdEncoding.DecodedLen(len(src))
        dst := make([]byte, maxlen)
        n, err := base32.StdEncoding.Decode(dst, []byte(src))
        fmt.Println(err != nil) // true
        fmt.Println(n <= maxlen) // true
        fmt.Println(n != len(origin)) // true
        fmt.Println(string(dst[0:n]) == "this ") // true

    }

    func main() {

        Example1()
        Example2()
        Example3()

    }
