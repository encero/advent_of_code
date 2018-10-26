package main

import (
	"fmt"
	"github.com/encero/advent_of_code/tools"
	"regexp"
	"strconv"
)

//                                 kd dec -37 if gm <= 9
var rex = regexp.MustCompile("^([a-z]+) (inc|dec) (-?[0-9]+) if ([a-z]+) ([<>=!]+) (-?[0-9]+)$")

type Instruction struct {
	Target string
	Action string
	Amount int64

	Source    string
	Condition string
	Compare   int64
}

func NewInstruction(s string) Instruction {
	matches := rex.FindStringSubmatch(s)

	amount, _ := strconv.ParseInt(matches[3], 10, 64)
	compare, _ := strconv.ParseInt(matches[6], 10, 64)

	return Instruction{
		Target: matches[1],
		Action: matches[2],
		Amount: amount,

		Source:    matches[4],
		Condition: matches[5],
		Compare:   compare,
	}
}

func main() {
	lines := tools.ReadLines("./day_8.txt")

	regs := make(map[string]int64)
	totalMax := int64(0)

	for _, l := range lines {
		ist := NewInstruction(l)

		source := regs[ist.Source]

		ok := false

		switch ist.Condition {
		case "==":
			ok = source == ist.Compare
		case "!=":
			ok = source != ist.Compare
		case ">=":
			ok = source >= ist.Compare
		case "<=":
			ok = source <= ist.Compare
		case "<":
			ok = source < ist.Compare
		case ">":
			ok = source > ist.Compare
		default:
			panic("unknown condition " + ist.Condition)
		}

		if ok {
			switch ist.Action {
			case "inc":
				regs[ist.Target] += ist.Amount
			case "dec":
				regs[ist.Target] -= ist.Amount
			default:
				panic("unknown action " + ist.Action)
			}
		}

		if totalMax < regs[ist.Target] {
			totalMax = regs[ist.Target]
		}
	}

	max := int64(0)

	for _, v := range regs {
		if v > max {
			max = v
		}
	}

	fmt.Println("max register", max)
	fmt.Println("total max register", totalMax)
}
