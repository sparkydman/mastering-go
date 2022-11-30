package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Firstname string
	Lastname  string
	Tel       []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func WriteJson(output *os.File, key interface{}) error {
	encoder := json.NewEncoder(output)
	if err := encoder.Encode(key); err != nil {
		return err
	}
	return nil
}

func main() {
	jsonRecord := Record{
		Firstname: "Mihalis",
		Lastname: "Tsoukalos",
		Tel: []Telephone{
			{  true,  "1234-567" },
			{  true,  "1234-abcd" },
			{  false,  "abcc-567" },
		},
}

if err := WriteJson(os.Stdout, jsonRecord); err != nil {
	fmt.Println(err)
	os.Exit(1)
}

}