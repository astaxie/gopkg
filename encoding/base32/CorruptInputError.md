## func (e CorruptInputError) Error() string

参数列表：

无

返回值：

- 错误描述

功能说明：

返回错误描述

代码实例：

```go
package examples

import (
    "fmt"
    "encoding/base32"
)

func ExampleCorruptInputError1() {

    // 不合法的 base32 串
    src := "ORUGS4ZANF@ZSAYJAORSXG5BAON2HE2LOM4XA===="
    maxlen := base32.StdEncoding.DecodedLen(len(src))
    dst := make([]byte, maxlen)
    _, err := base32.StdEncoding.Decode(dst, []byte(src))
    if err != nil {
        fmt.Println(err.Error())
    }

    // Output:
    // illegal base32 data at input byte 10

}
```
