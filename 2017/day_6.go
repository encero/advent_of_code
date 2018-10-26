package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/encero/advent_of_code/tools"
	"strings"
)

func main() {
	mems := tools.StringsToInts(strings.Split("0	5	10	0	11	14	13	4	11	8	8	7	1	4	12	11", "\t"))

	count := len(mems)
	total := memsSum(mems)

	cycles := 0
	loopCount := 0

	before := make(map[string]bool)
	before[signature(mems)] = true

	matchPatter := ""

	for {
		pointer := findMax(mems)

		toDistribute := mems[pointer]
		mems[pointer] = 0
		for i := int64(0); i < toDistribute; i++ {
			pointer += 1

			if pointer >= count {
				pointer = 0
			}

			mems[pointer] += 1
		}

		if total != memsSum(mems) {
			panic("sum changed")
		}

		cycles += 1

		sig := signature(mems)

		if matchPatter == "" {
			if _, ok := before[sig]; ok {
				matchPatter = sig
			}

			before[sig] = true
		} else {
			loopCount += 1
			if sig == matchPatter {
				break
			}
		}

	}

	fmt.Println("same pattern found after", cycles, "cycles")
	fmt.Println("loop length is", loopCount, "cycles")

}

func signature(mems []int64) string {
	buff := &bytes.Buffer{}

	for _, v := range mems {
		binary.Write(buff, binary.LittleEndian, v)
		buff.Write([]byte{0x0})
	}

	return buff.String()
}

func findMax(mems []int64) int {
	max := mems[0]

	for _, v := range mems {
		if max < v {
			max = v
		}
	}

	for i, v := range mems {
		if v == max {
			return i
		}
	}

	panic("max not found")
}

func memsSum(mems []int64) int64 {
	sum := int64(0)
	for _, v := range mems {
		sum += v
	}

	return sum
}

