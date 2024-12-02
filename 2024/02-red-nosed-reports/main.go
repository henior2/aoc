package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafe(report []int) bool {
	prev := 0
	increasing := false

	for i, v := range report {
		if i == 0 {
			prev = v
			continue
		}

		if i == 1 {
			increasing = v > prev
		}

		if dist := abs(v - prev); dist < 1 || dist > 3 {
			return false
		}

		if (increasing && v < prev) || (!increasing && v > prev) {
			return false
		}
	}

	return true
}

func main() {
	data, err := os.ReadFile("02.in")
	check(err)
	
	lines := strings.Split(string(data), "\n")
	
	safeReports := 0

	for _, line := range lines {
		line = strings.Trim(line, " ")

		report := make([]int, strings.Count(line, " "))

		for j, num := range strings.Split(line, " ") {
			report[j], err = strconv.Atoi(num)
			check(err)
		}

		if isSafe(report) {
			safeReports++
		}
	}

	fmt.Println("Safe reports:", safeReports)
}