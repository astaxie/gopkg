# json
功能介绍:
主要用于处理JSON数据，序列化和反序列化等。
相关文档可以参考，"JSON and Go"  [http://golang.org/doc/articles/json_and_go.html]( http://golang.org/doc/articles/json_and_go.html)

函数列表:

- [func Compact(dst *bytes.Buffer, src []byte) error](Compact.md)
- [func HTMLEscape(dst *bytes.Buffer, src []byte)](HTMLEscape.md)
- [func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error](Indent.md)
- [func Marshal(v interface{}) ([]byte, error)](Marshal.md)
- [func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)](MarshalIndent.md)
- [func Unmarshal(data []byte, v interface{}) error](Unmarshal.md)
- type Decoder
 - [func NewDecoder(r io.Reader) *Decoder](NewDecoder.md)
 - [func (dec *Decoder) Decode(v interface{}) error](Decode.md)
- type Encoder
 - [func NewEncoder(w io.Writer) *Encoder](NewEncoder.md)
 - [func (enc *Encoder) Encode(v interface{}) error](Encoder.md)
- type InvalidUTF8Error
 - func (e *InvalidUTF8Error) Error() string
- type InvalidUnmarshalError
 - func (e *InvalidUnmarshalError) Error() string
- type Marshaler
- type MarshalerError
 - func (e *MarshalerError) Error() string
- type RawMessage
 - func (m *RawMessage) MarshalJSON() ([]byte, error)
 - func (m *RawMessage) UnmarshalJSON(data []byte) error
- type SyntaxError
 - func (e *SyntaxError) Error() string
- type UnmarshalFieldError
 - func (e *UnmarshalFieldError) Error() string
- type UnmarshalTypeError
 - func (e *UnmarshalTypeError) Error() string
- type Unmarshaler
- type UnsupportedTypeError
 - func (e *UnsupportedTypeError) Error() string
- type UnsupportedValueError
 - func (e *UnsupportedValueError) Error() string