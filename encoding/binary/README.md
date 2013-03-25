# encoding/binary函数列表

- func PutUvarint(buf []byte, x uint64) int
- func PutVarint(buf []byte, x int64) int
- func Read(r io.Reader, order ByteOrder, data interface{}) error
- func ReadUvarint(r io.ByteReader) (uint64, error)
- func ReadVarint(r io.ByteReader) (int64, error)
- func Size(v interface{}) int
- func Uvarint(buf []byte) (uint64, int)
- func Varint(buf []byte) (int64, int)
- func Write(w io.Writer, order ByteOrder, data interface{}) error
