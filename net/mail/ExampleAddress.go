package main

import (
	"net/mail"
	"fmt"
)

func main() {
	addr := mail.Address{"Jim Green", "noreply@noreply.com"}
	fmt.Println(addr.String())
	// Output: =?utf-8?q?Jim_Green?= <noreply@noreply.com>
	addr = mail.Address{"中文", "noreply@noreply.com"}
	fmt.Println(addr.String())
	// Output: =?utf-8?q?=E4=B8=AD=E6=96=87?= <noreply@noreply.com>
}
