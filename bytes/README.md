# bytes

bytes包实现了用于管理字节切片的函数，类似于strings包。

### Constants

	const MinRead = 512
	
### Variables

	var ErrTooLarge = errors.New("bytes.Buffer: too large")

- [func Compare(a, b []byte) int](Compare.md)
- [func Contains(b, subslice []byte) bool](Contains.md)
- [func Count(s, sep []byte) int](Count.md)
- [func Equal(a, b []byte) bool](Equal.md)
- [func EqualFold(s, t []byte) bool](EqualFold.md)
- [func Fields(s []byte) [][]byte](Fields.md)
- [func FieldsFunc(s []byte, f func(rune) bool) [][]byte](FieldsFunc.md)
- [func HasPrefix(s, prefix []byte) bool](HasPrefix.md)
- [func HasSuffix(s, suffix []byte) bool](HasSuffix.md)
- [func Index(s, sep []byte) int](Index.md)
- [func IndexAny(s []byte, chars string) int](IndexAny.md)
- [func IndexByte(s []byte, c byte) int](IndexByte.md)
- [func IndexFunc(s []byte, f func(r rune) bool) int](IndexFunc.md)
- [func IndexRune(s []byte, r rune) int](IndexRune.md)
- [func Join(a [][]byte, sep []byte) []byte](Join.md)
- [func LastIndex(s, sep []byte) int](LastIndex.md)
- [func LastIndexAny(s []byte, chars string) int](LastIndexAny.md)
- [func LastIndexFunc(s []byte, f func(r rune) bool) int](LastIndexFunc.md)
- [func Map(mapping func(r rune) rune, s []byte) []byte](Map.md)
- [func Repeat(b []byte, count int) []byte](Repeat.md)
- [func Replace(s, old, new []byte, n int) []byte](Replace.md)
- [func Runes(s []byte) []rune](Runes.md)
- [func Split(s, sep []byte) [][]byte](Split.md)
- [func SplitAfter(s, sep []byte) [][]byte](SplitAfter.md)
- [func SplitAfterN(s, sep []byte, n int) [][]byte](SplitAfterN.md)
- [func SplitN(s, sep []byte, n int) [][]byte](SplitN.md)
- [func Title(s []byte) []byte](Title.md)
- [func ToLower(s []byte) []byte](ToLower.md)
- [func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte](ToLowerSpecial.md)
- [func ToTitle(s []byte) []byte](ToTitle.md)
- [func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte](ToTitleSpecial.md)
- [func ToUpper(s []byte) []byte](ToUpper.md)
- [func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte](ToUpperSpecial.md)
- [func Trim(s []byte, cutset string) []byte](Trim.md)
- [func TrimFunc(s []byte, f func(r rune) bool) []byte](TrimFunc.md)
- [func TrimLeft(s []byte, cutset string) []byte](TrimLeft.md)
- [func TrimLeftFunc(s []byte, f func(r rune) bool) []byte](TrimLeftFunc.md)
- [func TrimRight(s []byte, cutset string) []byte](TrimRight.md)
- [func TrimRightFunc(s []byte, f func(r rune) bool) []byte](TrimRightFunc.md)
- [func TrimSpace(s []byte) []byte](TrimSpace.md)

### type Buffer

	type Buffer struct {
		// contains filtered or unexported fields
    }

- [func NewBuffer(buf []byte) *Buffer](NewBuffer.md)
- [func NewBufferString(s string) *Buffer](NewBufferString.md)
- [func (b *Buffer) Bytes() []byte](Buffer_Bytes.md)
- [func (b *Buffer) Len() int](Buffer_Len.md)
- [func (b *Buffer) Next(n int) []byte](Buffer_Next.md)
- [func (b *Buffer) Read(p []byte) (n int, err error)](Buffer_Read.md)
- [func (b *Buffer) ReadByte() (c byte, err error)](Buffer_ReadByte.md)
- [func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)](Buffer_ReadBytes.md)
- [func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)](Buffer_ReadFrom.md)
- [func (b *Buffer) ReadRune() (r rune, size int, err error)](Buffer_ReadRune.md)
- [func (b *Buffer) ReadString(delim byte) (line string, err error)](Buffer_ReadString.md)
- [func (b *Buffer) Reset()](Buffer_Reset.md)
- [func (b *Buffer) String() string](Buffer_String.md)
- [func (b *Buffer) Truncate(n int)](Buffer_Truncate.md)
- [func (b *Buffer) UnreadByte() error](Buffer_UnreadByte.md)
- [func (b *Buffer) UnreadRune() error](Buffer_UnreadRune.md)
- [func (b *Buffer) Write(p []byte) (n int, err error)](Buffer_Write.md)
- [func (b *Buffer) WriteByte(c byte) error](Buffer_WriteByte.md)
- [func (b *Buffer) WriteRune(r rune) (n int, err error)](Buffer_WriteRune.md)
- [func (b *Buffer) WriteString(s string) (n int, err error)](Buffer_WriteString.md)
- [func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)](Buffer_WriteTo.md)

### type Reader

	type Reader struct {
		// contains filtered or unexported fields
	}
	
- [func NewReader(b []byte) *Reader](NewReader.md)
- [func (r *Reader) Len() int](Reader_Len.md)
- [func (r *Reader) Read(b []byte) (n int, err error)](Reader_Read.md)
- [func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)](Reader_ReadAt.md)
- [func (r *Reader) ReadByte() (b byte, err error)](Reader_ReadByte.md)
- [func (r *Reader) ReadRune() (ch rune, size int, err error)](Reader_ReadRune.md)
- [func (r *Reader) Seek(offset int64, whence int) (int64, error)](Reader_Seek.md)
- [func (r *Reader) UnreadByte() error](Reader_UnreadByte.md)
- [func (r *Reader) UnreadRune() error](Reader_UnreadRune.md)
