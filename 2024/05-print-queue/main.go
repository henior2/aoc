package main

import (
	"fmt"
	"os"
	"strings"

	"sync"
	"sync/atomic"
)

var rules [][2]int

func checkUpdate(update []int,  wg *sync.WaitGroup, midSum *atomic.Uint32) {
	defer wg.Done()

	pos := make(map[int]int)
	for i, v := range update {
		pos[v] = i
	}



	fmt.Println(update)
}

func main() {
	filename := "eg.in"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	rulesEnd := 0
	for lines[rulesEnd] != "" {
		rulesEnd++
	}

	rules := make([][2]int, rulesEnd)
	
	for i, r := range lines[:rulesEnd] {
		var a, b int
		n, err := fmt.Sscanf(r, "%d|%d", &a, &b)

		if n != 2 || err != nil {
			panic(err) // again, it's dumb
		}

		rules[i] = [2]int{a, b}
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
		checkUpdate(update, &wg, &midSum)
	}

	wg.Wait()

	fmt.Printf("Middle sums: %d\n", midSum.Load())

}
