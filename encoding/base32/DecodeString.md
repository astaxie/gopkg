## func (enc *Encoding) DecodeString(s string) ([]byte, error)

参数列表：

- s 要进行 base32 解码的字符串

返回值：

- 经过 base32 解码后的字符串切片
- 可能的错误

功能说明：

对传入的字符串进行 base323 解码

代码实例：

    package main

    import (
        "fmt"
        "encoding/base32"
    )

    func Example1() {

        fmt.Println("---=== Example1 ===---")

        src := "ORUGS4ZANFZSAYJAORSXG5BAON2HE2LOM4XA===="
        dst, _ := base32.StdEncoding.DecodeString(src)
        fmt.Println(string(dst) == "this is a test string.") // true

        dst, _ = base32.StdEncoding.DecodeString("")
        fmt.Println(string(dst) == "") // true

        dst, _ = base32.StdEncoding.DecodeString("MY======")
        fmt.Println(string(dst) == "f") // true

        dst, _ = base32.StdEncoding.DecodeString("MZXQ====")
        fmt.Println(string(dst) == "fo") // true

        dst, _ = base32.StdEncoding.DecodeString("MZXW6===")
        fmt.Println(string(dst) == "foo") // true

        dst, _ = base32.StdEncoding.DecodeString("MZXW6YQ=")
        fmt.Println(string(dst) == "foob") // true

        dst, _ = base32.StdEncoding.DecodeString("MZXW6YTB")
        fmt.Println(string(dst) == "fooba") // true

        dst, _ = base32.StdEncoding.DecodeString("MZXW6YTBOI======")
        fmt.Println(string(dst) == "foobar") // true
    }

    func main() {

        Example1()

    }
