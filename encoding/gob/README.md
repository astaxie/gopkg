# encoding/gob


gob is Go's own format for serializing and deserializing program data in binary format 

函数列表
- [func NewDecoder(r io.Reader) *Decoder](Decode_Encode.md)
- [func (dec *Decoder) Decode(e interface{}) error](Decode_Encode.md)
- [func (dec *Decoder) DecodeValue(v reflect.Value) error](DecodeValue.md)
- [func NewEncoder(w io.Writer) *Encoder](Decode_Encode.md)
- [func (enc *Encoder) Encode(e interface{}) error](Decode_Encode.md)
- [func (enc *Encoder) EncodeValue(value reflect.Value) error](EncodeValue.md)


结构体：
- [type Decoder struct](struct.md)
- [type Encoder struct](struct.md)

