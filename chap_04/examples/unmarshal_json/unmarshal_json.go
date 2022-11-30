package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide a valid filename")
		os.Exit(1)
	}
	filename, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println("can't read file",err)
		os.Exit(1)
	}
	var rec map[string]interface{}

	err = json.Unmarshal(filename, &rec)
	 if err != nil {
		fmt.Println(err)
		os.Exit(1)
	 }

	// fmt.Println(string(filename))
	fmt.Println(rec)
}