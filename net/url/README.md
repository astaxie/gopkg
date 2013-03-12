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
- func ParseQuery(query string) (m Values, err error)
