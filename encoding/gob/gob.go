package main

import (
	"encoding/gob"
	"fmt"
	//"log"
	//bufio"
	//io"
	//io/ioutil"
	"os"
	//"strings"
	"bytes"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func List_VCard(vc VCard) {
	fmt.Printf("FirstName = %s\n", vc.FirstName)
	fmt.Printf("Lastname = %s\n", vc.LastName)
	for i := 0; i < len(vc.Addresses); i++ {
		fmt.Printf("Type = %s\n", vc.Addresses[i].Type)
		fmt.Printf("City = %s\n", vc.Addresses[i].City)
		fmt.Printf("Country = %s\n", vc.Addresses[i].Country)
	}
	fmt.Printf("Remark = %s\n", vc.Remark)
}

func Encoding_vcard() error {
	//list one
	pa := &Address{Type: "private", City: "Aartselaar", Country: "Belgium"}
	wa := &Address{Type: "work", City: "Boom", Country: "Belgium"}
	vc := VCard{"dnv aps", "sn", []*Address{pa, wa}, "none"}

	//list two
	vc_1 := VCard{"Jimes", "Lester", []*Address{pa, wa}, "Mark"}
	vc_2 := VCard{"杨", "万荣", []*Address{pa, wa}, "Mark"}

	var vc_arry []VCard = []VCard{vc, vc_1, vc_2}

	//using an encoder
	file, e := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	if e != nil {
		return e
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	err := enc.Encode(vc_arry)

	if err != nil {
		return err
	}

	return nil

}

func Decoding_vcard() error {
	var vc []VCard //
	file, err := os.OpenFile("vcard.gob", os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(file)

	err = dec.Decode(&vc)
	if err != nil {
		return err
	}
	for i := 0; i < len(vc); i++ {
		fmt.Printf("index = %d\n", i+1)
		List_VCard(vc[i])
	}

	//fmt.Println(vc)

	return nil
}

/*
gob.CommonType
gob.Decoder
gob.Encoder
gob.GobEncoder
gob.NewDecoder()
gob.NewEncoder()
gob.Register()
gob.RegisterName()

*/

func Decode_t() {

	type dect struct {
		Id   int
		Name string
	}

	var buf bytes.Buffer
	/*
		job := dect{Id: 1, Name: "Tom"}
		enc := gob.NewEncoder(os.Stdout)
		if err := enc.Encode(&job); err != nil {
			fmt.Printf("failed encoing to writer %v\n", err)
		}

		dec := gob.NewDecoder(
		if err := dec.Decode(&job); err != nil {
			fmt.Printf("failed decoding to writer %v\n", err)
		}
	*/
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	pa := []dect{{1, "Tom"}, {2, "jimes"}, {3, "杨万荣"}}
	err := enc.Encode(pa)
	if err != nil {
		fmt.Printf("Encodeing error : %v\n", err)
	}

	var get []dect
	err = dec.Decode(&get)
	if err != nil {
		fmt.Printf("Decoding error: %v\n", err)
	}
	fmt.Println(get)

}
func main() {
	Encoding_vcard()
	Decoding_vcard()
	Decode_t()
}
