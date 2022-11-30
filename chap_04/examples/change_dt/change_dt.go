package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Enter one log path")
		os.Exit(1)
	}

	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Can't open log file", err)
		os.Exit(1)
	}
	defer file.Close()

	read := bufio.NewReader(file)
	notMatch := 0

	for {
		line, err := read.ReadString('\n')
		if err == io.EOF {
			break
		}else if err != nil {
			fmt.Println("can not read log line", err)
		}

		r1 := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)
		if r1.MatchString(line){
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line,match[1], newFormat, 1))
			}else{
				notMatch++
			}
			continue
		}
		r2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\]
.*`)
		if r2.MatchString(line){
			match := r2.FindStringSubmatch(line)
			d2, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d2.Format(time.Stamp)
				fmt.Print(strings.Replace(line,match[1], newFormat, 1))
			}else{
				notMatch++
			}
			continue
		}
	}
	fmt.Println(notMatch, "lines did not match!")

}