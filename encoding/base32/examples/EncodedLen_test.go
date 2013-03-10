package examples

import (
    "fmt"
    "encoding/base32"
)

func ExampleEncodedLen1() {

    src := []byte("this is a test string.")

    // 编码需要的缓冲区长度
    l := base32.StdEncoding.EncodedLen(len(src))

    dst := base32.StdEncoding.EncodeToString(src)
    fmt.Println(len(dst) == l)

    // Output:
    // true
}
