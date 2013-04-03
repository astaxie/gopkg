# func SendMail

```go
func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
```

net/smtp 里最重要的函数, 一般发送邮件任务都可以胜任.

```go
package main

import (
        "log"
        "net/mail"
        "encoding/base64"
        "net/smtp"
        "fmt"
        "strings"
)

func encodeRFC2047(String string) string{
        // use mail's rfc2047 to encode any string
        addr := mail.Address{String, ""}
        return strings.Trim(addr.String(), " <>")
}


func main() {
        // Set up authentication information.

        smtpServer := "smtp.163.com"
        auth := smtp.PlainAuth(
                "",
                "fledna@163.com",
                "PASSWORD",
                smtpServer,
        )

        from := mail.Address{"监控中心", "fledna@163.com"}
        to := mail.Address{"收件人", "haskell@renren.com"}
        title := "当前时段统计报表"

        body := "报表内容一切正常";

        header := make(map[string]string)
        header["From"] = from.String()
        header["To"] = to.String()
        header["Subject"] = encodeRFC2047(title)
        header["MIME-Version"] = "1.0"
        header["Content-Type"] = "text/plain; charset=\"utf-8\""
        header["Content-Transfer-Encoding"] = "base64"

        message := ""
        for k, v := range header {
                message += fmt.Sprintf("%s: %s\r\n", k, v)
        }
        message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))
        println(message)
        // Connect to the server, authenticate, set the sender and recipient,
        // and send the email all in one step.
        err := smtp.SendMail(
                smtpServer + ":25",
                auth,
                from.Address,
                []string{to.Address},
                []byte(message),
        )
        if err != nil {
                log.Fatal(err)
        }
}
```
