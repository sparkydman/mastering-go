package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	jsonRecord := Record{
		Firstname: "Mihalis",
		Lastname:  "Tsoukalos",
		Tel: []Telephone{
			{true, "1234-567"},
			{true, "1234-abcd"},
			{false, "abcc-567"},
		},
	}

	record, err := json.Marshal(&jsonRecord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(record))
}