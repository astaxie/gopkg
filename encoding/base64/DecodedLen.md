## func (enc *Encoding) DecodedLen(n int) int

参数列表：

- n 要进行 base64 解码的字节数

返回值：

- base64 解码输出需要的最大字节数

功能说明：

计算 base64 解码输出需要的最大字节数

代码实例：

```go
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
```
