func NewDecoder(r io.Reader) *Decoder

参数列表:

- r Reader对象

返回值:

- *Decoder 指向Decoder的指针

功能说明:

这个函数主要是给r创建一个decoder实例

代码实例:

    package main

	import (
		"encoding/json"
		"fmt"
		"io"
		"log"
		"strings"
	)
	
	func main() {
		const jsonStream = `
			{"Name": "Ed", "Text": "Knock knock."}
			{"Name": "Sam", "Text": "Who's there?"}
			{"Name": "Ed", "Text": "Go fmt."}
			{"Name": "Sam", "Text": "Go fmt who?"}
			{"Name": "Ed", "Text": "Go fmt yourself!"}
		`
		type Message struct {
			Name, Text string
		}
		dec := json.NewDecoder(strings.NewReader(jsonStream))
		for {
			var m Message
			if err := dec.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s: %s\n", m.Name, m.Text)
		}
	}





显示结果:

	Ed: Knock knock.
	Sam: Who's there?
	Ed: Go fmt.
	Sam: Go fmt who?
	Ed: Go fmt yourself!