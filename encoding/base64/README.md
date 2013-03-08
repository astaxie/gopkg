# encoding/base64 包函数列表

- Variables
- [func NewDecoder(enc *Encoding, r io.Reader) io.Reader](NewDecoder.md)
- [func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser](NewEncoder.md)
- type CorruptInputError
-     func (e CorruptInputError) Error() string
- type Encoding
-     [func NewEncoding(encoder string) *Encoding](NewEncoding.md)
-     [func (enc *Encoding) Decode(dst, src []byte) (n int, err error)](Decode.md)
-     [func (enc *Encoding) DecodeString(s string) ([]byte, error)](DecodeString.md)
-     [func (enc *Encoding) DecodedLen(n int) int](DecodedLen.md)
-     [func (enc *Encoding) Encode(dst, src []byte)](Encode.md)
-     [func (enc *Encoding) EncodeToString(src []byte) string](EncodeToString.md)
-     [func (enc *Encoding) EncodedLen(n int) int](EncodedLen.md)
