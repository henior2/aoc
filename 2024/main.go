package main

import (
	"fmt"

	"flag"

	"github.com/henior2/aoc/2024/day01"
)

type dayFunc func(part int, args []string)

func main() {
	dayModules := []dayFunc{
		day1.Run,
	}

	var day, part int
	flag.IntVar(&day, "day", -1, "The day of which solution to run (-1 for last)")
	flag.IntVar(&part, "part", 2, "The part of the problem (1 or 2)")

	flag.Parse()

	if day == -1 {
		day = len(dayModules)
	}

	if day < 1 || day > len(dayModules) {
		panic(fmt.Sprintf("Error: No such day: %d", day))
	}

	if part < 1 || part > 2 {
		panic(fmt.Sprintf("Error: No such part: %d\n Part must be an integer between 1 and 2", part))
	}

	dayModules[day-1](part, flag.Args())
}