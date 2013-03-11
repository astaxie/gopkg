#  bufio包函数列表

## type ReadWriter

- [func NewReadWriter(r *Reader, w *Writer) *ReadWriter](ReadWriter.md)

## type Reader

- [func NewReader(rd io.Reader) *Reader](NewReader.md)
- [func NewReaderSize(rd io.Reader, size int) *Reader](NewReaderSize.md)
- [*func (b *Reader) Buffered() int](Reader_Buffered.md)
- [*func (b *Reader) Peek(n int)](Reader_Peek.md)
- [*func (b *Reader) Read(p []byte) (n int, err error)](Reader_Read.md)

## type Writer

- [func NewWriter(wr io.Writer) *Writer](NewWriter.md)
- [func NewWriterSize(wr io.Writer, size int) *Writer](NewWriterSize.md)
- xx
