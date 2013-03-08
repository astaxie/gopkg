## type Message

```go
type Message struct {
    Header Header
    Body   io.Reader
}
```

Message 是邮件消息的 Golang 表示，代表一个已经解析了的邮件对象。

函数列表

- func ReadMessage(r io.Reader) (msg *Message, err error)

```go
package main

import (
	"net/mail"
	"bytes"
	"io/ioutil"
	"fmt"
)

func main() {
	rawMessage := `From: Feather <dummy@dummy.com>
To: noreply@noreply.com, noreply2@noreply.com
Subject: =?utf-8?B?5rWL6K+V5Y+R6YCB?=
Date: Fri, 21 May 2010 08:54:49 +0800
Message-ID: <m3fx1mt6p2.fsf@dummy.com>
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: plain

Hello Golang Mail`

	buf := bytes.NewBuffer([]byte(rawMessage))
	// parse mail message, buf may from net or somewhere.
	msg, err := mail.ReadMessage(buf)
	if err != nil {
		panic(err)
	}
	// mail body
	body, _ := ioutil.ReadAll(msg.Body)
	fmt.Printf("Mail Body:\n%s\n", body)

	fmt.Println("To Addresses:")
	printAddrs(msg, "to")

	fmt.Println("From Address:")
	printAddrs(msg, "From")

	// auto convert to golang time.Time
	date, _ := msg.Header.Date()
	fmt.Println("Date:", date)

	// mime header will ignore case
	fmt.Println("mime version:", msg.Header.Get("mime-version"))
}

func printAddrs(msg *mail.Message, filedName string) {
	addrs, err := msg.Header.AddressList(filedName)
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		fmt.Println("Addr:", addr)
	}
}
```

示例中有几个需要注意的地方

- 邮件头的 mime 域不区分大小写，适用于 Get()、 AddressList() 。
- RFC 2047 编码的字符串不一定会被解码，所以 Subject 不能被还原。
  （希望未来版本能完善）
- 由于邮件服务比较混乱，客户端对标准的实现参差不齐，所以 AddressList()
  的解析不一定成功
- 如果无法信任 net/mail， 建议自己解析

发送邮件内容请参考 net/smtp。
