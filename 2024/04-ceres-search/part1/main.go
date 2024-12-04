package main

import (
	"fmt"
	"os"

	"strings"
)

func findWord(word string, i, j int, grid []string) int {
	found := 0

	canGoRight := j < len(grid[i]) - 4

	if canGoRight && grid[i][j:j+4] == word {
		found++
	}

	if i+3 < len(grid) {
		vertical := []byte{grid[i][j], grid[i+1][j], grid[i+2][j], grid[i+3][j]}	
		if string(vertical) == word  {
			found++
		}
	
		if canGoRight {
			diagonalDown := []byte{grid[i][j], grid[i+1][j+1], grid[i+2][j+2], grid[i+3][j+3]}
			if string(diagonalDown) == word {
				found++
			}
		}
	}

	if canGoRight && i >= 3 {
		diagonalUp := []byte{grid[i][j], grid[i-1][j+1], grid[i-2][j+2], grid[i-3][j+3]}
		if string(diagonalUp) == word {
			found++
		}
	}

	return found
}

func main() {
	filename := "../04.in"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	crossword := strings.Split(string(data), "\n")

	found := 0


	for i, line := range crossword {
		for j, c := range line {

			if c != 'X' && c != 'S' {
				continue
			}

			found += findWord("XMAS", i, j, crossword)
			found += findWord("SAMX", i, j, crossword)

		}
	}
	fmt.Printf("Found: %d\n", found)

}
