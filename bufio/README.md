#  bufio包函数列表

### Variables

	var (
		ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
		ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
		ErrBufferFull        = errors.New("bufio: buffer full")
		ErrNegativeCount     = errors.New("bufio: negative count")
	)

### type ReadWriter

	type ReadWriter struct {
		*Reader
		*Writer
	}
	
- [func NewReadWriter(r *Reader, w *Writer) *ReadWriter](ReadWriter.md)

### type Reader
	
	type Reader struct {
		// contains filtered or unexported fields
	}

- [func NewReader(rd io.Reader) *Reader](NewReader.md)
- [func NewReaderSize(rd io.Reader, size int) *Reader](NewReaderSize.md)
- [func (b *Reader) Buffered() int](Reader_Buffered.md)
- [func (b *Reader) Peek(n int) ([]byte, error)](Reader_Peek.md)
- [func (b *Reader) Read(p []byte) (n int, err error)](Reader_Read.md)
- [func (b *Reader) ReadByte() (c byte, err error)](Reader_ReadByte.md)
- [func (b *Reader) ReadBytes(delim byte) (line []byte, err error)](Reader_ReadBytes.md)
- [func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)](Reader_ReadLine.md)
- [func (b *Reader) ReadRune() (r rune, size int, err error)](Reader_ReadRune.md)
- [func (b *Reader) ReadSlice(delim byte) (line []byte, err error)](Reader_ReadSlice.md)
- [func (b *Reader) ReadString(delim byte) (line string, err error)](Reader_ReadString.md)
- [func (b *Reader) UnreadByte() error](Reader_UnreadByte.md)
- [func (b *Reader) UnreadRune() error](Reader_UnreadRune.md)

### type Writer

	type Writer struct {
		// contains filtered or unexported fields
	}

- [func NewWriter(wr io.Writer) *Writer](NewWriter.md)
- [func NewWriterSize(wr io.Writer, size int) *Writer](NewWriterSize.md)
- [func (b *Writer) Available() int](Writer_Available.md)
- [func (b *Writer) Buffered() int](Writer_Buffered.md)
- [func (b *Writer) Flush() error](Writer_Flush.md)
- [func (b *Writer) Write(p []byte) (nn int, err error)](Writer_Write.md)
- [func (b *Writer) WriteByte(c byte) error](Writer_WriteByte.md)
- [func (b *Writer) WriteRune(r rune) (size int, err error)](Writer_WriteRune.md)
- [func (b *Writer) WriteString(s string) (int, error)](Writer_WriteString.md)
