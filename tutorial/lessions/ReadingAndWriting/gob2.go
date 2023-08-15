package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Address2 struct {
	Type    string
	City    string
	Country string
}

// The following is an example of a mixture of readable and binary data, as you will see when you try to read it in a text editor:
type VCard2 struct {
	FirstName string
	LastName  string
	Addresses []*Address2
	Remark    string
}

func main() {
	pa := &Address2{"private", "Aartselaar", "Belgium"}
	wa := &Address2{"work", "Boom", "Belgium"}
	vc := VCard2{"Jan", "Kersschot", []*Address2{pa, wa}, "none"}

	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("D:/LearnGolang/basic/lessions/ReadingAndWriting/vcard2.gob", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}
