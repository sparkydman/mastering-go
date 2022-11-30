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

func LoadFromJson(filename string, key interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(key); err != nil {
		return err
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Json file required")
		os.Exit(1)
	}

	var record Record

	if err := LoadFromJson(os.Args[1], &record); err != nil {
		fmt.Println("Can not load json file", err)
		os.Exit(1)
	}

	fmt.Println(record)
}