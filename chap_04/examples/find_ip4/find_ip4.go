package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIp(input string) string {
	ipParttern := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	ipGrammr := ipParttern + "\\." + ipParttern + "\\." + ipParttern + "\\." + ipParttern
	matchIp := regexp.MustCompile(ipGrammr)

	return	matchIp.FindString(input)
}

func main() {
	args := os.Args
	if len(args) < 2{
		fmt.Printf("usage: %s logFile\n", filepath.Base(args[0]))
 		os.Exit(1)
	}

	for _, filename := range args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("can't open log file", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}else if err != nil {
			fmt.Println("Can't read log line", err)
			continue
		}

		ipMatch := findIp(line)
		trial := net.ParseIP(ipMatch)
		if trial.To4() == nil {
			continue
		}
		fmt.Println(ipMatch)
	}
}