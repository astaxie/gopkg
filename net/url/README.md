# 包名:net/url
# 概术:
  Package url parses URLs and implements query escaping. 
# 函数列表
- [func QueryEscape(s string) string](QueryEscape.md)
- [func QueryUnescape(s string) (string, error)](QueryUnescape.md)
- type Error
  - func (e *Error) Error() string
- type EscapeError
  - func (e EscapeError) Error() string
- type URL
  - [func Parse(rawurl string) (url *URL, err error)](Parse.md)
  - func ParseRequestURI(rawurl string) (url *URL, err error)
  - func (u *URL) IsAbs() bool
  - func (u *URL) Parse(ref string) (*URL, error)
  - func (u *URL) Query() Values
  - func (u *URL) RequestURI() string
  - func (u *URL) ResolveReference(ref *URL) *URL
  - func (u *URL) String() string
- type Userinfo
  - func User(username string) *Userinfo
  - func UserPassword(username, password string) *Userinfo
  - func (u *Userinfo) Password() (string, bool)
  - func (u *Userinfo) String() string
  - func (u *Userinfo) Username() string
- [type Values](Values.md)
  - [func ParseQuery(query string) (m Values, err error)](QueryUnescape.md)
  - [func (Values) Add](Add.md)
  - [func (v Values) Del(key string)](Del.md)
  - [func (v Values) Encode() string](Encode.md)
  - [func (v Values) Get(key string) string](Get.md)
  - [func (v Values) Set(key, value string)](Set.md)
