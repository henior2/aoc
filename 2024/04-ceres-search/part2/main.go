package main

import (
	"fmt"
	"os"

	"strings"
)

func isX(i, j int, grid []string) bool {

	ul, ur := grid[i-1][j-1], grid[i-1][j+1] // up left, up right
	dl, dr := grid[i+1][j-1], grid[i+1][j+1] // down left, down right

	for _, corner := range []byte{ul, ur, dl, dr} {
		if corner != 'M' && corner != 'S' { // all corners need to be M or S
			return false
		}
	}

	horizontalLines := ul == ur && dl == dr && ul != dl
	verticalLines := ul == dl && ur == dr && ul != ur

	return horizontalLines || verticalLines
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


	for i := 1; i < len(crossword) - 1; i++ {
		for j := 1; j < len(crossword[i]) - 1; j++ {
			if crossword[i][j] != 'A' {
				continue
			}
			
			if isX(i, j, crossword) {
				found++
			}
		}
	}
	fmt.Printf("Found: %d\n", found)

}
