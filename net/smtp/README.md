# net/smtp

RFC 5321 (SMTP, Simple Mail Transfer Protocol, 简单邮件传输协议) 的 Golang 实现。想发邮件，你就找对了。

类型列表

- [Auth](Auth.md) - smtp 用户认证方法
- [Client](Client.md) - smtp 客户端实现
- ServerInfo

ServerInfo 只用于传参, 为简单 struct, 没有成员函数.

函数列表

- func [SendMail](SendMail.md)(addr string, a Auth, from string, to []string, msg []byte) error

一般简单的发送邮件任务请直接是用 SendMail() 函数.
