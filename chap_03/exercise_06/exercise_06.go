package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Plese enter a date - in \"dd MMM yyy\" format")
	}

	enteredDate := os.Args[1]

	d, err := time.Parse("02 Jan 2006", enteredDate)
	if err != nil {
		log.Fatalf("Can't parse date: %v", err)
	}

	fmt.Printf("Full date: %v\n", d)
	fmt.Printf("Day: %v\t Month: %v\t Year: %v\t", d.Day(), d.Month(), d.Year()) 
}