package main

import (
	"fmt"
	"github.com/encero/advent_of_code/tools"
)

func main() {
	lines := tools.ReadLines("./day_5.txt")

	jumps := tools.StringsToInts(lines)

	pointer := int64(0)
	steps := 0
	for {
		if pointer >= int64(len(jumps)) {
			break
		}

		offset := jumps[pointer]
		steps += 1

		if offset >= 3 {
			jumps[pointer] -= 1
		} else {
			jumps[pointer] += 1
		}

		pointer += offset
	}

	fmt.Println("escaped after", steps, "steps")

}
