## func (enc *Encoding) EncodeToString(src []byte) string

参数列表

- src 要进行 base64 编码的字符串切片

返回值：

- 经过 base64 编码后的字符串

功能说明：

对传入的字符串切片进行 base64 编码

代码实例：

    package main

    import (
        "fmt"
        "encoding/base64"
    )

    func main() {

        src := []byte("this is a test string.")
        dst := base64.StdEncoding.EncodeToString(src)
        fmt.Println(dst == "dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg==") // true

        fmt.Println(base64.StdEncoding.EncodeToString([]byte("")) == "") // true
        fmt.Println(base64.StdEncoding.EncodeToString([]byte("f")) == "Zg==") // true
        fmt.Println(base64.StdEncoding.EncodeToString([]byte("fo")) == "Zm8=") // true
        fmt.Println(base64.StdEncoding.EncodeToString([]byte("foo")) == "Zm9v") // true
        fmt.Println(base64.StdEncoding.EncodeToString([]byte("foob")) == "Zm9vYg==") // true
        fmt.Println(base64.StdEncoding.EncodeToString([]byte("fooba")) == "Zm9vYmE=") // true
        fmt.Println(base64.StdEncoding.EncodeToString([]byte("foobar")) == "Zm9vYmFy") // true
    }
