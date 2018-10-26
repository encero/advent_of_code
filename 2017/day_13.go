package main

import (
	"fmt"
	"github.com/encero/advent_of_code/tools"
	"strconv"
	"strings"
)

func main() {
	lines := tools.ReadLines("./day_13.txt")

	layers := make(map[int64]int64)

	last := int64(0)

	for _, line := range lines {
		splt := strings.Split(line, ":")

		layer, _ := strconv.ParseInt(splt[0], 10, 64)
		depth, _ := strconv.ParseInt(strings.Trim(splt[1], " "), 10, 64)

		layers[layer] = depth

		if layer > last {
			last = layer
		}
	}

	offset := int64(0)

	for {
		offset += 1

		caught := false
		for layer := int64(0); layer <= last; layer ++ {
			depth, ok := layers[layer]

			if !ok {
				continue
			}

			if Position(depth, layer+offset) == 0 {
				caught = true
			}
		}

		if offset%1000 == 0 {
			fmt.Println("off", offset)
		}

		if ! caught {
			break
		}
	}

	fmt.Println("offset", offset)
}

func Position(depth, time int64) int64 {
	dir := int64(1)
	cur := int64(0)

	cycle := (depth * 2) - 2

	cycles := time / (cycle)

	time = time - (cycles * cycle)

	for i := int64(0); i < time; i++ {
		cur += dir

		if cur == 0 || (cur+1) == depth {
			dir = -1 * dir
		}

	}

	return cur
}

// 0 1 2 3 4 3 2 1
func Position2(depth, time int64) int64 {
	dir := int64(1)
	cur := int64(0)

	//if time == 0 {
	//	return 0
	//}

	for i := int64(0); i < time; i++ {
		cur += dir

		if cur == 0 || (cur+1) == depth {
			dir = -1 * dir
		}

	}

	return cur
}
