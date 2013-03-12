# bytes

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

函数列表

- xxx1
- xxx2
