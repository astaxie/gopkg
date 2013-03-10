package examples

import (
    "fmt"
    "encoding/base64"
)

func ExampleDecodedLen1() {

    src := "dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg=="

    // 解码需要的缓冲区的最大长度
    l := base64.StdEncoding.DecodedLen(len(src))

    dst, _ := base64.StdEncoding.DecodeString(src)
    fmt.Println(len(dst) <= l)

    // Output:
    // true
}
