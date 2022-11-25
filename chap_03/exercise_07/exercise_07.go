package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Plese enter a time - in \"HH:mm:ss\" format")
	}

	enteredTime := os.Args[1]

	t, err := time.Parse("15:04:05", enteredTime)
	if err != nil {
		log.Fatalf("Can't parse Time: %v", err)
	}

	fmt.Printf("Full Time: %v\n", t)
	fmt.Printf("Hour: %v\t Minute: %v\t Second: %v\t", t.Hour(), t.Minute(), t.Second()) 
}