package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	defer input.Close()

	safeCount := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		reportSlice := strings.Split(scanner.Text(), " ")
		reportIntSlice, err := convReportToInts(reportSlice)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		test1 := allIncreasingOrDecreasing(reportIntSlice)
		test2 := diffIsSafe(reportIntSlice)
		if test1 && test2 {
			safeCount = safeCount + 1
		} else {
			fmt.Println("failed, testing removal of items")
			fmt.Println("original:", reportIntSlice)
			if testFailedReports(reportIntSlice) {
				fmt.Println("passed with removal")
				safeCount = safeCount + 1
			}
		}
	}
	fmt.Println(safeCount)

}

func convReportToInts(report []string) ([]int, error) {
	var intReport []int
	for _, v := range report {
		intV, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		intReport = append(intReport, intV)
	}

	return intReport, nil
}

func allIncreasingOrDecreasing(report []int) bool {
	var previous bool
	var current bool
	for i, _ := range report {
		if i == 0 {
			continue
		}
		if report[i] == report[i-1] {
			return false
		}
		current = report[i] > report[i-1]
		if i == 1 {
			previous = report[i] > report[i-1]
		}
		if current != previous {
			return false
		}
	}
	return true
}

func diffIsSafe(report []int) bool {
	for i, _ := range report {
		if i == 0 {
			continue
		}
		diff := math.Abs(float64(report[i] - report[i-1]))
		// fmt.Println(diff)
		if diff == 0 || diff > 3 {
			return false
		}
	}
	return true
}

func newSlice(i int, s []int) []int {
	ns := make([]int, len(s))
	copy(ns, s)
	return append(ns[:i], ns[i+1:]...)
}

func testFailedReports(report []int) bool {
	for i := range report {
		newReport := newSlice(i, report)
		test1 := allIncreasingOrDecreasing(newReport)
		test2 := diffIsSafe(newReport)
		if test1 && test2 {
			return true
		}
	}

	return false
}
