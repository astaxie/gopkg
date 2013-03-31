## type Auth

Auth 是一个接口类型, 用于表示 smtp 认证过程中的具体方法.

这里简单科普下, 具体内容请查阅 RFC. SMTP 中一般有三种认证方法, LOGIN, PLAIN, CRAM-MD5. net/smtp
中有后两种的实现, 缺少一个 LOGIN 实现. 一般国内邮件服务都支持前两种认证方法, CRAM-MD5 支持商不多.

```go
type Auth
    func CRAMMD5Auth(username, secret string) Auth
    func PlainAuth(identity, username, password, host string) Auth
```

PlainAuth 的使用例子可以参考 [SendMail](SendMail.md).

下面这个略复杂的例子里, 简单实现了一个 LOGIN 认证的 Auth 对象, 并发送邮件.

```go
package main

import (
        "log"
        "net/mail"
        "net/smtp"
        "fmt"
        "strings"
        "errors"
)

func encodeRFC2047(String string) string{
        // use mail's rfc2047 to encode any string
        addr := mail.Address{String, ""}
        return strings.Trim(addr.String(), " <>")
}

type loginAuth struct {
        username, password string
}

func LoginAuth(username, password string) smtp.Auth {
        return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
        if !server.TLS {
                log.Fatal("unencrypted connection")
        }
        return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
        if more {
                switch string(fromServer) {
                case "Username:":
                        return []byte(a.username), nil
                case "Password:":
                        return []byte(a.password), nil
                default:
                        return nil, errors.New("Unkown fromServer")
                }
        }
        return nil, nil
}

func main() {
        var err error
        // Set up authentication information.
        smtpServer := "smtp.qq.com:25"
        client, err := smtp.Dial(smtpServer)
        handleError(err)
        defer client.Quit()

        if ok, desc := client.Extension("AUTH"); ok {
                fmt.Println("wow! server support auth:", desc)
        }

        err = client.Auth(LoginAuth("username@qq.com", "PASSWORD"))
        handleError(err)

        handleError(client.Mail("username@qq.com"))
        handleError(client.Rcpt("fledna@foxmail.com"))
        writer, err := client.Data()
        handleError(err)
        writer.Write([]byte(`MIME-Version: 1.0
Subject: =?utf-8?q?=E5=BD=93=E5=89=8D=E6=97=B6=E6=AE=B5=E7=BB=9F=E8=AE=A1=E6=8A=A5=E8=A1=A8?=
Content-Type: text/plain; charset="utf-8"
From: =?utf-8?q?=E7=9B=91=E6=8E=A7=E4=B8=AD=E5=BF=83?= <username@qq.com>
Content-Transfer-Encoding: base64
To: =?utf-8?q?=E6=94=B6=E4=BB=B6=E4=BA=BA?= <fledna@foxmail.com>

5oql6KGo5YaF5a655LiA5YiH5q2j5bi4`))
        handleError(writer.Close())
}


func handleError(err error) {
        if err != nil {
                log.Fatal(err)
        }
}
```
