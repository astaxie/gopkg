## type Address

```golang
type Address struct {
    Name    string // 专有名称；可以为空。
    Address string // user@domain
}
```

Address 用于表示邮件中的一个 email 地址，具体格式请参考
RFC 5322，源代码中还包含一个 RFC 2047 的编码函数，
但是没有暴露出来。不过通过小技巧可以使用，请参考例子。

函数列表

- func (a *Address) String() string

邮件地址转换

```go
package main

import (
    "net/mail"
    "fmt"
)

func main() {
    addr := mail.Address{"Jim Green", "noreply@noreply.com"}
    fmt.Println(addr.String())
    // Output: =?utf-8?q?Jim_Green?= <noreply@noreply.com>
    addr = mail.Address{"中文", "noreply@noreply.com"}
    fmt.Println(addr.String())
    // Output: =?utf-8?q?=E4=B8=AD=E6=96=87?= <noreply@noreply.com>
}
```

RFC 2047 编码函数

```go
// import net/mail, strings
func EncodeRFC2047(String string) string{
    // use mail's rfc2047 to encode any string
    addr := mail.Address{String, ""}
    return strings.Trim(addr.String(), " <>")
}
```
