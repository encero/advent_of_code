package main

import (
	"fmt"
	"github.com/encero/advent_of_code/tools"
)

const stNone = "none"
const stGarbage = "garbage"

func main() {
	lines := tools.ReadLines("./day_9.txt")
	line := lines[0]

	state := stNone
	opened := 0
	negate := false

	score := 0
	garbage := 0

	for i := range line {
		ch := line[i:i+1]

		fmt.Print(ch)

		switch state {
		case stNone:
			negate = false
			if ch == "{" {
				opened += 1
				score += opened
				fmt.Print(" open ", opened)
			}
			if ch == "}" {
				fmt.Print(" close ", opened)
				opened -= 1
			}
			if ch == "<" {
				state = stGarbage
				fmt.Print(" start garbage")
			}
		case stGarbage:
			if negate {
				negate = false
				fmt.Print(" negated")
			} else if ch == ">" {
				state = stNone
				fmt.Print(" stop garbage")
			} else if  ch == "!" {
				negate = true
				fmt.Print(" negating next char")
			} else {
				garbage += 1
				fmt.Print(" garbage")
			}
		}

		fmt.Println()
	}

	fmt.Println("total score", score, "total garbage", garbage)
}
