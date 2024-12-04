package main

import (
	"aoc-cli/day01"
	"aoc-cli/day02"
	"fmt"
	"os"
)

func main() {
	day := os.Args[1]
	switch day {
	case "day-one":
		fmt.Printf("Day 01 - Location ID distance for %v", os.Args[2])
		day01.Day1(os.Args[2])

	case "day-two":
		fmt.Printf("Day 02 - Report safety for %v", os.Args[2])
		day02.Day2(os.Args[2])
	}
}
