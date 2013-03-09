## type Header

```go
type Header map[string][]string
```

Header 用于表示邮件头中的键值对(Key-Value Pairs)。使用的时候尽量使用 Get() 和 AddressList() 方法操作，别作为 map 操作。

函数列表

- func (h Header) AddressList(key string) ([]*Address, error)
- func (h Header) Date() (time.Time, error)
- func (h Header) Get(key string) string

该类型一般用于解析，构建消息时由于它缺乏对应的方法，不够方便。相关例子请参考 [Message](Message.md)。

如果你熟悉 Golang 其他库的话，那么你会发现 textproto.MIMEHeader 和
mail.Header 以及 http.Header 其实都是 ``map[string][]string``。
所以部分缺失的方法可以用其他库代替，不过这不是好编程习惯。
