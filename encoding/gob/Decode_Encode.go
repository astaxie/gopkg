package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {

	type dect struct {
		Id   int
		Name string
	}

	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	err := enc.Encode(dect{2, "mary"})
	if err != nil {
		fmt.Printf("Encodeing error : %v\n", err)
	}

	var get dect
	err = dec.Decode(&get)
	if err != nil {
		fmt.Printf("Decoding error: %v\n", err)
	}
	fmt.Println(get)

}

//运行
//{2,mary}
