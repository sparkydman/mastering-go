package main

import (
	"fmt"
	"os"
	"strconv"
)

// Write a constant generator iota for the days of the week.

const (
	monday = 1 + iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

var weekday = map[int]string{
	monday: "MONDAY",
	tuesday: "TUESDAY",
	wednesday: "WEDNESDAY",
	thursday: "THURSDAY",
	friday: "FRIDAY",
	saturday: "SATURDAY",
	sunday: "SUNDAY",
}

func getDayOfWeek(day int) string{
	switch day{
	case monday:
		return	weekday[monday]
	case tuesday:
		return weekday[tuesday]
	case wednesday:
		return weekday[wednesday]
	case thursday:
		return weekday[thursday]
	case friday:
		return weekday[friday]
	case saturday:
		return weekday[saturday]
	case sunday:
		return weekday[sunday]
	default:
		return "Not valid weekday"
	}
}

func main() {
	if len(os.Args) != 2 {	
		fmt.Println("Please provide a day")
		return
	}

	day := os.Args[1]

	n, err := strconv.Atoi(day)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(getDayOfWeek(n))
}