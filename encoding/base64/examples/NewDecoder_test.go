package examples

import (
    "fmt"
    "strings"
    "encoding/base64"
)

// 从 string 取得输入
func ExampleNewDecoder1() {

    src := "dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg=="
    reader := strings.NewReader(src)

    dst := ""

    decoder := base64.NewDecoder(base64.StdEncoding, reader)
    // 使用一个很小的输出buffer，测试流式解码
    buf := make([]byte, 2)
    for {
        n, err := decoder.Read(buf)
        if err != nil {
            break
        }
        if n == 0 {
            break
        }
        dst += string(buf[0:n])
    }
    fmt.Print(dst)

    // Output:
    // this is a test string.

}
