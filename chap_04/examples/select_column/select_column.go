package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
Write a program that takes number of column and list of files path and loop through them and print the 15th word of the file
*/

func main() {
	args := os.Args;
	if len(args) < 2 {
		fmt.Printf("usage: selectColumn column <file> <file> [...<file>]\n")
		os.Exit(1)
	}

	column, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid number of columns", err)
		return
	}

	if column < 0 {
		fmt.Printf("Invalid number of columns %v", column)
	}

	for _, file := range args[2:]{
		fmt.Println("*****************READING FILE*************", file)
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Can't open file", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)

		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}else{
			fmt.Println("Can't read file", err)
		}

		words := strings.Fields(line)
		if len(words) >= column{
			fmt.Println(words[column-1])
		}
	}
}