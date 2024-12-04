package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day3(filepath string, conditionals bool) {

	f, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(fmt.Errorf("error opening file: %v", err))
	}

	var res [][]string
	if conditionals {
		fmt.Println("Calculating with conditionals")
		res = findMultiplicationsWithCond(string(f))
	} else {
		fmt.Println("Calculating without conditionals")
		res = findMultiplications(string(f))
	}
	total := 0
	enabled := true
	for _, v := range res {
		if v[0] == "don't()" {
			enabled = false
			continue
		} else if v[0] == "do()" {
			enabled = true
			continue
		}

		if enabled {
			ans, err := multiply(v[0])
			if err != nil {
				fmt.Println("error multiplying", err)
			}
			total = total + ans
		}
	}

	fmt.Println(total)
}

func findMultiplications(s string) [][]string {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllStringSubmatch(s, -1)

	return matches
}

func findMultiplicationsWithCond(s string) [][]string {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
	matches := r.FindAllStringSubmatch(s, -1)

	return matches
}

func multiply(s string) (int, error) {
	r := regexp.MustCompile(`(\d{1,3})`)
	matches := r.FindAllStringSubmatch(s, -1)
	n1, n1err := strconv.Atoi(matches[0][0])
	n2, n2err := strconv.Atoi(matches[1][0])
	if n1err != nil || n2err != nil {
		return 0, fmt.Errorf("error converting string to int")
	}

	return n1 * n2, nil
}
