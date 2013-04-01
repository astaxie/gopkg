# encoding/binary

此包实现了对数据与byte之间的转换，以及varint的编解码。

##函数列表

数据的byte序列化转换
- [func Read(r io.Reader, order ByteOrder, data interface{}) error](Read.md)
- [func Write(w io.Writer, order ByteOrder, data interface{}) error](Write.md)
- [func Size(v interface{}) int](Size.md)

uvarint和varint的编解码
- [func PutUvarint(buf []byte, x uint64) int](PutUvarint.md)
- [func PutVarint(buf []byte, x int64) int](PutVarint.md)
- [func Uvarint(buf []byte) (uint64, int)](Uvarint.md)
- [func Varint(buf []byte) (int64, int)](Varint.md)
- [func ReadUvarint(r io.ByteReader) (uint64, error)](ReadUvarint.md)
- [func ReadVarint(r io.ByteReader) (int64, error)](ReadVarint.md)

##结构体
- type ByteOrder：可以定义自己的字节序结构，用于序列化和反序列化数据。
