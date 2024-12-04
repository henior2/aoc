package main

import (
	"fmt"
	"os"

	"strconv"
	"strings"

	"sync"
	"sync/atomic"
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

func isSafe(report []int, skip int) bool {
	prev := 0
	increasing := false

	hadSkipped := 0	

	for i, v := range report {
		if i == skip {
			hadSkipped = 1
			continue
		}

		if i - hadSkipped == 0 { // nothing to compare with
			prev = v
			continue
		}
		
		if i - hadSkipped == 1 { // first pair increasing?
			increasing = v > prev
		}
		
		if dist := abs(v - prev); dist < 1 || dist > 3 {
			return false
		}
		
		if (increasing && v < prev) || (!increasing && v > prev) {
			return false
		}

		prev = v
	}
	
	return true
}

func dampener(report []int, wg *sync.WaitGroup, safeReports *atomic.Uint64) {
	defer wg.Done()
	
	for i := range report {
		if isSafe(report, i) {
			// fmt.Println(report, i)

			safeReports.Add(1)
			return
		}
	}
}

func main() {
	filename := "../02.in"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	data, err := os.ReadFile(filename)
	check(err)

	lines := strings.Split(string(data), "\n")

	var wg sync.WaitGroup
	var safeReports atomic.Uint64

	for _, line := range lines {
		line = strings.Trim(line, " ")
		line = strings.TrimSuffix(line, "\r") // Windows moment

		report := make([]int, strings.Count(line, " ")+1)

		for j, num := range strings.Split(line, " ") {
			report[j], err = strconv.Atoi(num)
			check(err)
		}

		// if isSafe(report) {
		// 	safeReports++
		// }

		wg.Add(1)
		go dampener(report, &wg, &safeReports)
	}

	 

	fmt.Println("Safe reports:", safeReports.Load())
}
