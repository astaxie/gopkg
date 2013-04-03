# encoding/base32 包

实现了 RFC 4648 定义的 base32 编码

### 变量

RFC 4648 定义的标准 base32 编码
```go
var StdEncoding = NewEncoding(encodeStd)
```

RFC 4648 里定义的 “Extended Hex Alphabet”, 一般用于 DNS
```go
var HexEncoding = NewEncoding(encodeHex)
```

### 函数

- [func NewDecoder(enc *Encoding, r io.Reader) io.Reader](NewDecoder.md)
- [func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser](NewEncoder.md)

### CorruptInputError 类型

- [func (e CorruptInputError) Error() string](CorruptInputError.md)

### Encoding 类型

- [func NewEncoding(encoder string) *Encoding](NewEncoding.md)
- [func (enc *Encoding) Decode(dst, src []byte) (n int, err error)](Decode.md)
- [func (enc *Encoding) DecodeString(s string) ([]byte, error)](DecodeString.md)
- [func (enc *Encoding) DecodedLen(n int) int](DecodedLen.md)
- [func (enc *Encoding) Encode(dst, src []byte)](Encode.md)
- [func (enc *Encoding) EncodeToString(src []byte) string](EncodeToString.md)
- [func (enc *Encoding) EncodedLen(n int) int](EncodedLen.md)
