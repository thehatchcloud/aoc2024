package main

import (
	"aoc-cli/day01"
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "day-one" {
		fmt.Printf("Day 01 - Location ID distance against %v", os.Args[2])
		day01.Day1(os.Args[2])
	}
}
