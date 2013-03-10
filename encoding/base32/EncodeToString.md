## func (enc *Encoding) EncodeToString(src []byte) string

参数列表：

- src 要进行 base32 编码的字符串切片

返回值：

- 经过 base32 编码后的字符串

功能说明：

对传入的字符串切片进行 base32 编码

代码实例：

    package main

    import (
        "fmt"
        "encoding/base32"
    )

    func Example1() {

        fmt.Println("---=== Example1 ===---")

        src := []byte("this is a test string.")
        dst := base32.StdEncoding.EncodeToString(src)
        fmt.Println(dst == "ORUGS4ZANFZSAYJAORSXG5BAON2HE2LOM4XA====") // true

        fmt.Println(base32.StdEncoding.EncodeToString([]byte("")) == "") // true
        fmt.Println(base32.StdEncoding.EncodeToString([]byte("f")) == "MY======") // true
        fmt.Println(base32.StdEncoding.EncodeToString([]byte("fo")) == "MZXQ====") // true
        fmt.Println(base32.StdEncoding.EncodeToString([]byte("foo")) == "MZXW6===") // true
        fmt.Println(base32.StdEncoding.EncodeToString([]byte("foob")) == "MZXW6YQ=") // true
        fmt.Println(base32.StdEncoding.EncodeToString([]byte("fooba")) == "MZXW6YTB") // true
        fmt.Println(base32.StdEncoding.EncodeToString([]byte("foobar")) == "MZXW6YTBOI======") // true
    }

    func main() {

        Example1()

    }