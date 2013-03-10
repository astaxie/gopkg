## func (enc *Encoding) Encode(dst, src []byte)

参数列表：

- dst 编码缓冲区
- src 要编码的字符串切片

返回值：

无

功能说明：

对输入的字符串进行 base32 编码

代码实例：

```go
package examples

import (
    "fmt"
    "encoding/base32"
)

func ExampleEncode1() {

    // 要编码的字符串
    src := "this is a test string."
    // 计算编码缓冲区需要的长度
    l := base32.StdEncoding.EncodedLen(len(src))

    // 创建足够长度的缓冲区
    dst := make([]byte, l)

    // 编码
    base32.StdEncoding.Encode(dst, []byte(src))

    fmt.Println(string(dst) == "ORUGS4ZANFZSAYJAORSXG5BAON2HE2LOM4XA====")
    fmt.Println(l == len(dst))
    fmt.Println(l%8 == 0)

    // Output:
    // true
    // true
    // true

}

func ExampleEncode2() {

    // 编码缓冲区需要的长度会被填充到8的倍数
    for ii := 0; ii < 20; ii++ {
        l := base32.StdEncoding.EncodedLen(ii)
        fmt.Println(l, l%8 == 0) // true
    }

    // Output:
    // 0 true
    // 8 true
    // 8 true
    // 8 true
    // 8 true
    // 8 true
    // 16 true
    // 16 true
    // 16 true
    // 16 true
    // 16 true
    // 24 true
    // 24 true
    // 24 true
    // 24 true
    // 24 true
    // 32 true
    // 32 true
    // 32 true
    // 32 true

}
```
