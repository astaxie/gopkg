package examples

import (
    "fmt"
    "encoding/base32"
)

// 我们的编码表，32 字节
const encodeTest = "--------------------------------"

func ExampleNewEncoding1() {

    enc := base32.NewEncoding(encodeTest)

    src := "this is a test string."
    dst := enc.EncodeToString([]byte(src))

    // 最后不足8字节的会用"="补全
    fmt.Println(dst)
    fmt.Println(len(dst) % 8 == 0)

    // Output:
    // ------------------------------------====
    // true

}
