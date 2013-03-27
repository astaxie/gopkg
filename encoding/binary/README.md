# encoding/binary函数列表

- func Read(r io.Reader, order ByteOrder, data interface{}) error
- func Write(w io.Writer, order ByteOrder, data interface{}) error
- func Size(v interface{}) int
- [func PutUvarint(buf []byte, x uint64) int](PutUvarint.md)
- [func PutVarint(buf []byte, x int64) int](PutVarint.md)
- [func Uvarint(buf []byte) (uint64, int)](Uvarint.md)
- [func Varint(buf []byte) (int64, int)](Varint.md)
- [func ReadUvarint(r io.ByteReader) (uint64, error)](ReadUvarint.md)
- [func ReadVarint(r io.ByteReader) (int64, error)](ReadVarint.md)
