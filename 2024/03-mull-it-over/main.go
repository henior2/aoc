package main

import (
	"fmt"
	"os"

	"regexp"

	"cmp"

	pq "github.com/emirpasic/gods/v2/queues/priorityqueue"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}


type Command int
const (
	Mul Command = iota
	Do
	Dont
)

type Instruction struct {
	start int
	end int
	command Command
}

func addCommand(pattern, source string, command Command, q *pq.Queue[Instruction]) error {
	r, err := regexp.Compile(pattern)

	if err != nil {
		return err
	}

	for _, m := range r.FindAllStringIndex(source, -1) {
		q.Enqueue(Instruction{m[0], m[1], command})
	}

	return nil
}

func byStart(a, b Instruction) int {
	return cmp.Compare(a.start, b.start)
}

func main() {
	filename := "03.in"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	data_b, err := os.ReadFile(filename)
	check(err)

	data := string(data_b)

	insts := pq.NewWith(byStart)

	
	check( addCommand(`mul\(\d{1,3},\d{1,3}\)`, data, Mul, insts) )

	// comment out for part 1
	check( addCommand(`do\(\)`, data, Do, insts) )
	check( addCommand(`don't\(\)`, data, Dont, insts) )

	result := 0
	doing := true

	for !insts.Empty() {
		i, _ := insts.Dequeue()

		switch i.command {
		case Do:
			doing = true

		case Dont:
			doing = false

		case Mul:
			if !doing {
				continue
			}
			
			var a, b int
			fmt.Sscanf(data[i.start:i.end], "mul(%d,%d)", &a, &b)

			result += a * b
		}
	}

	fmt.Printf("Sum: %d\n", result)

}