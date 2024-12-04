package day1

import (
	"fmt"
	"os"
	"strings"
	"slices"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Run(part int, args []string) {
	filename := "test/01.in"
	if len(args) > 1 {
		filename = args[1]
	}

	data, err := os.ReadFile(filename)
	if (err != nil) {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	left := make([]int, len(lines))
	right := make([]int, len(lines))

	rightOccur := make(map[int]int)

	for i, line := range lines {
		n, err := fmt.Sscanf(line, "%d %d", &left[i], &right[i])
		
		if n != 2 || err != nil {
			panic(err) // this is dumb cause it will panic(nil) on n != 2 but whatever
		}

		rightOccur[right[i]]++
	}

	slices.Sort(left)
	slices.Sort(right)

	result := 0
	similarity := 0

	for i := range len(lines) {
		result += abs(left[i] - right[i])

		similarity += left[i] * rightOccur[left[i]]
	}
	
	fmt.Printf("Distance: %d\nSimilarity: %d\n", result, similarity)
	
}