package main

import (
	"aoc-cli/day01"
	"aoc-cli/day02"
	"aoc-cli/day03"
	"fmt"
	"os"
)

func main() {

	switch os.Args[1] {
	case "day-one":
		fmt.Printf("Day 01 - Location ID distance for %v", os.Args[2])
		day01.Day1(os.Args[2])

	case "day-two":
		fmt.Printf("Day 02 - Report safety for %v", os.Args[2])
		day02.Day2(os.Args[2])

	case "day-three":
		fmt.Printf("Day 03 - Cleaning multipliers for %v\n", os.Args[2])
		if len(os.Args) == 4 {
			fmt.Println(os.Args)
			if os.Args[3] == "conditionals" {
				day03.Day3(os.Args[2], true)
			}
		} else {
			day03.Day3(os.Args[2], false)
		}
	}
}
