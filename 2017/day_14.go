package main

import (
	"bytes"
	"fmt"
)

type Direction int

const (
	up Direction = iota
	left
	right
	down
)

func main() {

	grid := make([][]int, 128)

	for i := range grid {
		grid[i] = make([]int, 128)
	}

	count := 0
	for i := 0; i < 128; i++ {
		str := fmt.Sprint("uugsqrei-", i)
		//str := fmt.Sprint("flqrgnkx-", i)
		hs := hash(str)

		for j := range hs {
			if hs[j:j+1] == "1" {
				count += 1
				grid[i][j] = 1
			} else {
				grid[i][j] = 0
			}
		}
	}

	groups := 0
	for {
		x, y := find(grid)

		if x == -1 {
			break
		}

		groups += 1

		traverse(grid, x, y)
	}

	fmt.Println("count", count, "groups", groups)
}

type Pair struct {
	X int
	Y int
}

func find(grid [][]int) (int, int) {
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 1 {
				return x, y
			}
		}
	}

	return -1, -1
}

func traverse(grid [][]int, x, y int) {
	list := make([]Pair, 0)

	list = append(list, Pair{x, y})

	for i := 0; i < len(list); i++ {
		x := list[i].X
		y := list[i].Y

		if grid[x][y] != 1 {
			continue
		}

		grid[x][y] = 2

		for _, d := range []Direction{up, down, left, right} {
			switch d {
			case up:
				if y > 0 && grid[x][y-1] == 1 {
					list = append(list, Pair{x, y - 1})
				}
			case down:
				if y < 127 && grid[x][y+1] == 1 {
					list = append(list, Pair{x, y + 1})
				}
			case left:
				if x > 0 && grid[x-1][y] == 1 {
					list = append(list, Pair{x - 1, y})
				}
			case right:
				if x < 127 && grid[x+1][y] == 1 {
					list = append(list, Pair{x + 1, y})
				}
			}
		}
	}
}

func hash(input string) string {
	arr := make([]int, 256)
	for i := range arr {
		arr[i] = i
	}

	for _, v := range []int{17, 31, 73, 47, 23} {
		input += string(rune(v))
	}

	pointer := 0
	skip := 0
	for k := 0; k < 64; k++ {
		for _, v := range input {
			reverse(arr, pointer, int(v))

			pointer += int(v) + skip

			pointer = wrap(pointer, len(arr))
			skip += 1
		}
	}

	reduced := make([]int, 16)

	for i := range reduced {
		start := arr[i*16]
		for j := 1; j < 16; j++ {
			start ^= arr[(i*16)+j]
		}

		reduced[i] = start
	}

	out := &bytes.Buffer{}
	for _, v := range reduced {
		fmt.Fprintf(out, "%08b", uint8(v))
	}

	return out.String()
}

func reverse(arr []int, start, length int) {
	if length > (len(arr) - 1) {
		panic("len(arr) > length")
	}

	for i := (length / 2) - 1; i >= 0; i-- {
		a := wrap(start+length-i-1, len(arr))
		b := wrap(start+i, len(arr))

		tmp := arr[a]
		arr[a] = arr[b]
		arr[b] = tmp
	}
}

func wrap(x, max int) int {
	if x >= max {
		x = x - max

		return wrap(x, max)
	}

	return x
}
