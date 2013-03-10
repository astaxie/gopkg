package examples

import (
    "fmt"
    "encoding/base32"
)

func ExampleDecodedLen1() {

    src := "ORUGS4ZANFZSAYJAORSXG5BAON2HE2LOM4XA===="

    // 解码需要的缓冲区的最大长度
    l := base32.StdEncoding.DecodedLen(len(src))

    dst, _ := base32.StdEncoding.DecodeString(src)
    fmt.Println(len(dst) <= l)

    // Output:
    // true
}
