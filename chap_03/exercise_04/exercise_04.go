package main

import "fmt"

type power4 int

const (
	p4_0 power4 = 1 << iota * 1 << iota
	p4_1
	p4_2
	p4_3
)

func main() {
	fmt.Printf("4^%d = %d\n", 0, p4_0)
	fmt.Printf("4^%d = %d\n", 1, p4_1)
	fmt.Printf("4^%d = %d\n", 2, p4_2)
	fmt.Printf("4^%d = %d\n", 3, p4_3)
}