package main

import (
	"fmt"
	"os"
	"runtime"
	"slices"
	"strings"

	"sync"
	"sync/atomic"
)

const PART = 2

type rulesMap map[int] map[int]bool

type posMap map[int]int
func (pos posMap) cmpPages(a int, b int) int {
	posA, foundA := pos[a]
	posB, foundB := pos[b]

	if !foundA || !foundB {
		return 1
	}
	
	if posA < posB {
		return 1
	} else {
		return -1
	}
}

func checkUpdate(update []int, rules rulesMap, wg *sync.WaitGroup, midSum *atomic.Uint32) {
	defer wg.Done()

	pos := make(posMap)
	for i, v := range update {
		pos[v] = i
	}

	part2sorted := false

	for a, r := range rules {
		for b := range r {
			if pos.cmpPages(a, b) == 1 {
				continue
			}

			if PART == 2 {
				fmt.Print(update, " -> ")
				slices.SortFunc(update, pos.cmpPages)
				fmt.Println(update)
				part2sorted = true
				break
			}

			return
		}
		if part2sorted {
			break
		}
	}
	
	if PART == 1 || part2sorted {
		midSum.Add(uint32(update[len(update) / 2]))
	}

	// fmt.Println(update, "\n", pos, "\n")
}

func main() {
	filename := "05.in"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	newline := "\n"
	if runtime.GOOS == "windows" {
		newline = "\r\n"
	}

	lines := strings.Split(string(data), newline)

	rulesEnd := 0
	for lines[rulesEnd] != "" {
		rulesEnd++
	}

	rules := make(rulesMap)

	for _, r := range lines[:rulesEnd] {
		var a, b int
		n, err := fmt.Sscanf(r, "%d|%d", &a, &b)

		if n != 2 || err != nil {
			panic(err) // again, it's dumb
		}

		if rules[a] == nil {
			rules[a] = make(map[int]bool)
		}

		rules[a][b] = true
	}

	var wg sync.WaitGroup
	var midSum atomic.Uint32

	for _, u := range lines[rulesEnd+1:] {
		// u = strings.Trim(u, " ")
		// u = strings.TrimSuffix(u, "\r") // again, Windows

		update := make([]int, strings.Count(u, ",")+1)

		for j, num := range strings.Split(u, ",") {
			n, err := fmt.Sscanf(num, "%d", &update[j])
			if n != 1 || err != nil {
				panic(err)
			}
		}

		wg.Add(1)
		go checkUpdate(update, rules, &wg, &midSum)
	}

	wg.Wait()

	fmt.Printf("Middle sums: %d\n", midSum.Load())

}
