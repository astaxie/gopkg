#func ParseRequestURI(rawurl string) (url *URL, err error)
[参数列表]：

- rawurl 网址

[返回值]：

- 

[功能说明]：
解析一个URL到结构的rawurl。假定在HTTP请求中收到rawurl，所以rawurl仅解释为绝对URI或绝对路径。假定rawurl字符串后缀有“＃”。 （Web浏览器剥离“＃”片，然后再发送到Web服务器的URL。）
[代码实例]
