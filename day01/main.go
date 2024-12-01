package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

// Create a list (slice) from each column
// Sort each list
// Get the distance (difference) between each list
// Add all of the distances up to get the total distance

func main() {
	lineOne, lineTwo, err := getTwoSlices("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(lineOne) != len(lineTwo) {
		fmt.Println("number slices don't match in length")
		os.Exit(2)
	}
	sort.Ints(lineOne)
	sort.Ints(lineTwo)

	total := 0
	for i := range lineOne {
		// fmt.Printf("line %v: %v,%v\n", i+1, lineOne[i], lineTwo[i])
		var lineTotal int
		if lineOne[i] > lineTwo[i] {
			lineTotal = lineOne[i] - lineTwo[i]
		} else {
			lineTotal = lineTwo[i] - lineOne[i]
		}
		// fmt.Println("lineTotal:", lineTotal)
		total = total + lineTotal
	}
	fmt.Println("Total:", total)
	totalSimilarity := 0
	for _, v := range lineOne {
		c := countIntsInSlice(v, lineTwo)
		fmt.Println(v, "appears", c, "times in lineTwo")
		lineSimilarity := v * c
		fmt.Println("line similarity:", lineSimilarity)
		totalSimilarity = totalSimilarity + lineSimilarity
		fmt.Println("total similarity:", totalSimilarity)
	}

	fmt.Println("Final Total Similarity:", totalSimilarity)

}

func getTwoSlices(filename string) ([]int, []int, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	first := []int{}
	second := []int{}

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		count++
		// fmt.Println("line:", scanner.Text())
		numGroups := regexp.MustCompile(`(\s*)(\d*)`)
		matches := numGroups.FindAllStringSubmatch(scanner.Text(), -1)
		// fmt.Printf("first  = %v|\n", matches[0][0])
		firstInt, err := strconv.Atoi(matches[0][0])
		if err != nil {
			return nil, nil, fmt.Errorf("line %v - error converting first string to in: %v", count, err)
		}
		first = append(first, firstInt)
		// fmt.Printf("second = %v|\n", matches[1][2])
		secondInt, err := strconv.Atoi(matches[1][2])
		if err != nil {
			return nil, nil, fmt.Errorf("line %v - error converting second string to in: %v", count, err)
		}
		second = append(second, secondInt)

	}

	return first, second, nil
}

// func similarityScore(s1 []int, s2 []int) {
// 	for i, v := range s1 {
// 		fmt.Println(i, v)
// 	}

// }

func countIntsInSlice(number int, s []int) int {
	count := 0
	for _, v := range s {
		if v == number {
			count++
		}
	}
	return count
}
