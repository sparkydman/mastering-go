package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

/*
The developed utility, which is called rw_xml.go, will read a JSON record from disk, make
a change to it, convert it to XML, and print it on screen. Then it will convert the XML data
into JSON. The related Go code will be presented in four parts.
*/

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

	j := json.NewDecoder(f)
	if err = j.Decode(key); err != nil {
		return err
	}

	return nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: provide a valid json file path")
		os.Exit(1)
	}

	var record Record
	if err := LoadFromJson(args[1], &record); err != nil {
		fmt.Println("Can't load json file", err)
		os.Exit(1)
	}
	fmt.Println("JSON: ",record)

	xmlData, _ := xml.MarshalIndent(record, "", "\t")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Println("\nxmlData: ", string(xmlData))

	data := &Record{}

	if err := xml.Unmarshal(xmlData, data); err != nil {
		fmt.Println("can't unmarshal from xml",err)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("can't marshal from json",err)
		return
	}

	if err = json.Unmarshal([]byte(result), &record); err != nil {
		fmt.Println("can't unmarshal from json",err)
		return
	}

	fmt.Println("\nJSON:",record)

}