# text/scanner包函数列表

- [Constants](Constants.md)
- [func TokenString(tok rune) string](TokenString.md)
- [type Position](Position.md)
    - [func (pos *Position) IsValid() bool](IsValid.md)
    - [func (pos Position) String() string](String.md)
- type Scanner
	Scanner用于从io.Reader中读取Unicode字符或词法单元
    - [func (s *Scanner) Init(src io.Reader) *Scanner](Init.md)
    - [func (s *Scanner) Next() rune](Next.md)
    - [func (s *Scanner) Peek() rune](Peek.md)
    - [func (s *Scanner) Pos() (pos Position)](Pos.md)
    - [func (s *Scanner) Scan() rune](Scan.md)
    - [func (s *Scanner) TokenText() string](TokenText.md)

