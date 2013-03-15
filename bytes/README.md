# bytes

bytes包实现了用于管理字节切片的函数，类似于strings包。

### Constants

	const MinRead = 512
	
### Variables

	var ErrTooLarge = errors.New("bytes.Buffer: too large")

- [func Compare(a, b []byte) int]
- [func Contains(b, subslice []byte) bool]
- [func Count(s, sep []byte) int]
- [func Equal(a, b []byte) bool]
- [func EqualFold(s, t []byte) bool]
- [func Fields(s []byte) [][]byte]
- [func FieldsFunc(s []byte, f func(rune) bool) [][]byte]
- [func HasPrefix(s, prefix []byte) bool]
- [func HasSuffix(s, suffix []byte) bool]
- [func Index(s, sep []byte) int]
- [func IndexAny(s []byte, chars string) int]
- [func IndexByte(s []byte, c byte) int]
- [func IndexFunc(s []byte, f func(r rune) bool) int]
- [func IndexRune(s []byte, r rune) int]
- [func Join(a [][]byte, sep []byte) []byte]
- [func LastIndex(s, sep []byte) int]
- [func LastIndexAny(s []byte, chars string) int]
- [func LastIndexFunc(s []byte, f func(r rune) bool) int]
- [func Map(mapping func(r rune) rune, s []byte) []byte]
- [func Repeat(b []byte, count int) []byte]
- [func Replace(s, old, new []byte, n int) []byte]
- [func Runes(s []byte) []rune]
- [func Split(s, sep []byte) [][]byte]
- [func SplitAfter(s, sep []byte) [][]byte]
- [func SplitAfterN(s, sep []byte, n int) [][]byte]
- [func SplitN(s, sep []byte, n int) [][]byte]
- [func Title(s []byte) []byte]
- [func ToLower(s []byte) []byte]
- [func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte]
- [func ToTitle(s []byte) []byte]
- [func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte]
- [func ToUpper(s []byte) []byte]
- [func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte]
- [func Trim(s []byte, cutset string) []byte]
- [func TrimFunc(s []byte, f func(r rune) bool) []byte]
- [func TrimLeft(s []byte, cutset string) []byte]
- [func TrimLeftFunc(s []byte, f func(r rune) bool) []byte]
- [func TrimRight(s []byte, cutset string) []byte]
- [func TrimRightFunc(s []byte, f func(r rune) bool) []byte]
- [func TrimSpace(s []byte) []byte]

### type Buffer

	type Buffer struct {
		// contains filtered or unexported fields
    }

- [func NewBuffer(buf []byte) *Buffer]
- [func NewBufferString(s string) *Buffer]
- [func (b *Buffer) Bytes() []byte]
- [func (b *Buffer) Len() int]
- [func (b *Buffer) Next(n int) []byte]
- [func (b *Buffer) Read(p []byte) (n int, err error)]
- [func (b *Buffer) ReadByte() (c byte, err error)]
- [func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)]
- [func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)]
- [func (b *Buffer) ReadRune() (r rune, size int, err error)]
- [func (b *Buffer) ReadString(delim byte) (line string, err error)]
- [func (b *Buffer) Reset()]
- [func (b *Buffer) String() string]
- [func (b *Buffer) Truncate(n int)]
- [func (b *Buffer) UnreadByte() error]
- [func (b *Buffer) UnreadRune() error]
- [func (b *Buffer) Write(p []byte) (n int, err error)]
- [func (b *Buffer) WriteByte(c byte) error]
- [func (b *Buffer) WriteRune(r rune) (n int, err error)]
- [func (b *Buffer) WriteString(s string) (n int, err error)]
- [func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)]

### type Reader

	type Reader struct {
		// contains filtered or unexported fields
	}
	
- [func NewReader(b []byte) *Reader]
- [func (r *Reader) Len() int]
- [func (r *Reader) Read(b []byte) (n int, err error)]
- [func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)]
- [func (r *Reader) ReadByte() (b byte, err error)]
- [func (r *Reader) ReadRune() (ch rune, size int, err error)]
- [func (r *Reader) Seek(offset int64, whence int) (int64, error)]
- [func (r *Reader) UnreadByte() error]
- [func (r *Reader) UnreadRune() error]
