## func (enc *Encoding) EncodedLen(n int) int

参数列表：

- n 要进行 base32 编码的字节数

返回值：

- 经过 base32 编码后的字节数

功能说明：

计算编码输出的字节数

代码实例：

    package main

    import (
        "fmt"
        "encoding/base32"
    )

    func Example1() {

        fmt.Println("---=== Example1 ===---")

        src := []byte("this is a test string.")

        // 编码需要的缓冲区长度
        l := base32.StdEncoding.EncodedLen(len(src))

        dst := base32.StdEncoding.EncodeToString(src)
        fmt.Println(len(dst) == l) // true

    }

    func main() {

        Example1()

    }