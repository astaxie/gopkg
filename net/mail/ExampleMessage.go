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
