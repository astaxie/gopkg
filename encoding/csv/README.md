# encoding/csv包函数列表

- [func (e *ParseError) Error() string](Error.md)
- [func NewReader(r io.Reader) *Reader](NewReader.md)
- [func (r *Reader) Read() (record []string, err error)](Read.md)
- [func (r *Reader) ReadAll() (records [][]string, err error)](ReadAll.md)
- [func NewWriter(w io.Writer) *Writer](NewWriter.md)
- [func (w *Writer) Flush()](Flush.md)
- [func (w *Writer) Write(record []string) (err error)](Write.md)
- [func (w *Writer) WriteAll(records [][]string) (err error)](WriteAll.md)