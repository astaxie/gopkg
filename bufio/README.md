#  bufio包函数列表

### type ReadWriter

- [func NewReadWriter(r *Reader, w *Writer) *ReadWriter](ReadWriter.md)

### type Reader

- [func NewReader(rd io.Reader) *Reader](NewReader.md)
- [func NewReaderSize(rd io.Reader, size int) *Reader](NewReaderSize.md)
- [func (b *Reader) Buffered() int](Reader_Buffered.md)
- [func (b *Reader) Peek(n int)](Reader_Peek.md)
- [func (b *Reader) Read(p []byte) (n int, err error)](Reader_Read.md)
- [func (b *Reader) ReadByte() (c byte, err error)](Reader_ReadByte.md)
- [func (b *Reader) ReadBytes(delim byte) (line []byte, err error)]
- [func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)]
- [func (b *Reader) ReadRune() (r rune, size int, err error)]
- [func (b *Reader) ReadSlice(delim byte) (line []byte, err error)]
- [func (b *Reader) ReadString(delim byte) (line string, err error)]
- [func (b *Reader) UnreadByte() error]
- [func (b *Reader) UnreadRune() error]

### type Writer

- [func NewWriter(wr io.Writer) *Writer](NewWriter.md)
- [func NewWriterSize(wr io.Writer, size int) *Writer](NewWriterSize.md)
- [func (b *Writer) Available() int]
- [func (b *Writer) Buffered() int]
- [func (b *Writer) Flush() error]
- [func (b *Writer) Write(p []byte) (nn int, err error)]
- [func (b *Writer) WriteByte(c byte) error]
- [func (b *Writer) WriteRune(r rune) (size int, err error)]
- [func (b *Writer) WriteString(s string) (int, error)]
