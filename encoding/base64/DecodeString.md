## func (enc *Encoding) DecodeString(s string) ([]byte, error)

参数列表

- s 要进行 base64 解码的字符串

返回值：

- 经过 base64 解码后的字符串切片
- 可能的错误

功能说明：

对传入的字符串进行 base64 解码

代码实例：

    package main

    import (
        "fmt"
        "encoding/base64"
    )

    func main() {

        src := "dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg=="
        dst, _ := base64.StdEncoding.DecodeString(src)
        fmt.Println(string(dst) == "this is a test string.") // true

        dst, _ = base64.StdEncoding.DecodeString("Zg==")
        fmt.Println(string(dst) == "f") // true

        dst, _ = base64.StdEncoding.DecodeString("Zm8=")
        fmt.Println(string(dst) == "fo") // true

        dst, _ = base64.StdEncoding.DecodeString("Zm9v")
        fmt.Println(string(dst) == "foo") // true

        dst, _ = base64.StdEncoding.DecodeString("Zm9vYg==")
        fmt.Println(string(dst) == "foob") // true

        dst, _ = base64.StdEncoding.DecodeString("Zm9vYmE=")
        fmt.Println(string(dst) == "fooba") // true

        dst, _ = base64.StdEncoding.DecodeString("Zm9vYmFy")
        fmt.Println(string(dst) == "foobar") // true
    }
