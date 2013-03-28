# 包名  io

函数列表

- [func Copy(dst Writer, src Reader) (written int64, err error)](Copy.md)
- [func CopyN(dst Writer, src Reader, n int64) (written int64, err error)](CopyN.md)
- [func LimitReader(r Reader, n int64) Reader](LimitReader.md)
- [func (l *LimitedReader) Read(p []byte) (n int, err error)](LimitReader.Read.md)
- [func MultiReader(readers ...Reader) Reader](MultiReader.md)
- [func MultiWriter(writers ...Writer) Writer](MultiWriter.md)
- [func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader](NewSectionReade.mdr)
- [func Pipe() (*PipeReader, *PipeWriter)](Pipe.md)
- [func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)](ReadAtLeast.md)
- [func ReadFull(r Reader, buf []byte) (n int, err error)](ReadFull.md)
- [func (r *PipeReader) Close() error](PipeReader.Close.md)
- [func (r *PipeReader) CloseWithError(err error) error](PipeReader.CloseWithError.md)
- [func (r *PipeReader) Read(data []byte) (n int, err error)](PipeReader.Read.md)
- [func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)](SectionReader.ReadAt.md)
- [func (s *SectionReader) Read(p []byte) (n int, err error)](SectionReader.Read.md)
- [func (s *SectionReader) Seek(offset int64, whence int) (ret int64, err error)](SectionReader.Seek.md)
- [func (s *SectionReader) Size() int64](SectionReader.Size.md)
- [func TeeReader(r Reader, w Writer) Reader](TeeReader.Read.md)
- [func (w *PipeWriter) Close() error](PipeWriter.Close.md)
- [func (w *PipeWriter) CloseWithError(err error) error](PipeWriter.CloseWithError.md)
- [func (w *PipeWriter) Write(data []byte) (n int, err error)](PipeWriter.Write.md)
- [func WriteString(w Writer, s string) (n int, err error)](WriteString.md)
- 

