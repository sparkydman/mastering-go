package main

import "fmt"

var arr = [5]int{1, 2, 4, 5, 6}

func main() {
	var newMap = make(map[int]int, len(arr))

	for i := 0; i < len(arr); i++ {
		newMap[i] = arr[i]
	}

	fmt.Println(newMap)
}