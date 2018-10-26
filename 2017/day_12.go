package main

import (
	"fmt"
	"github.com/encero/advent_of_code/tools"
	"strconv"
	"strings"
)

type Pipes struct {
	Id        int64
	Connected []int64
	Visited   bool
}

func main() {
	lines := tools.ReadLines("./day_12.txt")

	pipes := make(map[int64]*Pipes)

	for _, v := range lines {
		first := strings.Split(v, "<->")
		progId, _ := strconv.ParseInt(strings.Trim(first[0], " "), 10, 64)

		ids := tools.StringsToInts(StringMap(strings.Split(first[1], ","), func(s string) string {
			return strings.Trim(s, " ")
		}))

		pipes[progId] = &Pipes{
			Id:        progId,
			Connected: ids,
		}
	}

	for i := range pipes {
		for j := range pipes[i].Connected {
			pipes[pipes[i].Connected[j]].Connected = append(pipes[pipes[i].Connected[j]].Connected, i)
		}
	}

	zero := pipes[0]

	fmt.Println("count of progs", traversePipes(pipes, zero))

	for _, v := range pipes {
		v.Visited = false // reset
	}

	groups := 0

	for id := range pipes {
		if pipes[id].Visited {
			continue
		}

		fmt.Println("traversing", id)

		groups += 1
		traversePipes(pipes, pipes[id])
	}

	fmt.Println("total groups", groups)
}

func traversePipes(pipes map[int64]*Pipes, pipe *Pipes) int {
	if pipe.Visited {
		return 0
	}

	pipe.Visited = true
	sum := 1

	for _, v := range pipe.Connected {
		sum += traversePipes(pipes, pipes[v])
	}

	return sum
}


func StringMap(arr []string, f func(string) string) []string {
	tmp := make([]string, len(arr))

	for i, s := range arr {
		tmp[i] = f(s)
	}

	return tmp
}
