# 包名 compress/flate

常量列表

    const (
        NoCompression = 0//不压缩
        BestSpeed     = 1//最快速度压缩
    
        BestCompression    = 9//最佳压缩比压缩
        DefaultCompression = -1//默认压缩
    )

函数列表

- [func NewReader(r io.Reader) io.ReadCloser](NewReader.md)
- [func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser](NewReaderDict.md)
- [func (e CorruptInputError) Error() string](CorruptInputError.Error.md)
- [func (e InternalError) Error() string](InternalError.Error.md)
- [func (e *ReadError) Error() string](ReadError.Error.md)
- [func (e *WriteError) Error() string](WriteError.Error.md)
- [func NewWriter(w io.Writer, level int) (*Writer, error)](NewWriter.md)
- [func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)](NewWriterDict.md)
- [func (w *Writer) Close() error](Writer.Close.md)
- [func (w *Writer) Flush() error](Writer.Flush.md)
- [func (w *Writer) Reset(dst io.Writer)](Writer.Reset.md)
- [func (w *Writer) Write(data []byte) (n int, err error)](Writer.Write.md)