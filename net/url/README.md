# 包名:net/url
# 概术:
  Package url parses URLs and implements query escaping. 
# 函数列表
- [func QueryEscape(s string) string](QueryEscape.md)
- [func QueryUnescape(s string) (string, error)](QueryUnescape.md)
- func Parse(rawurl string) (url *URL, err error)
- func ParseRequestURI(rawurl string) (url *URL, err error)
- func User(username string) *Userinfo
- func UserPassword(username, password string) *Userinfo
- [type Values](Values.md)
  - [func ParseQuery(query string) (m Values, err error)](QueryUnescape.md)
  - [func (Values) Add](Add.md)
  - [func (v Values) Del(key string)](Del.md)
  - [func (v Values) Encode() string](Encode.md)
  - [func (v Values) Get(key string) string](Get.md)
  - [func (v Values) Set(key, value string)](Set.md)
