package main

import (
	"fmt"
	"os"
)

const Target = 289326

func main() {
	//fmt.Println("number 1 total distance: ", totalDistance(1), "expected 0\n")
	//fmt.Println("number 12 total distance: ", totalDistance(12), "expected 3\n")
	//fmt.Println("number 23 total distance: ", totalDistance(23), "expected 2\n")
	//fmt.Println("number 1024 total distance: ", totalDistance(1024), "expected 31\n")
	//fmt.Println("number", Target, "total distance: ", totalDistance(Target))

	stressTest(Target)
}

func totalDistance(target int) int {
	level := 0
	edgeLength := 1
	currentPoints := 1
	spiralLength := 1

	for {
		level += 1                          // rise level one point
		edgeLength = edgeLength + 2         // edge length rise 2 point on each level
		spiralLength = (edgeLength * 4) - 4 // there is 4 times edge length minus corner overlaps
		currentPoints += spiralLength       // add spiral length

		if currentPoints >= target {
			fmt.Println("level", level, "edge len", edgeLength, "current points", currentPoints)

			totalDistance := level // first part of distance is level of spiral

			spiralPart := spiralLength / 4             // one side of spiral ( one point shorter due to overlaps )
			startPoint := currentPoints - spiralLength // start point of current spiral

			for partIndex := 1; partIndex <= 4; partIndex ++ {
				corner := startPoint + (spiralPart * partIndex)
				if corner >= target { // found part of spiral
					half := spiralPart / 2 // center of spiral part

					spiralPartHalfPoints := startPoint + (spiralPart * (partIndex - 1)) + half
					distanceFromCenter := abs(target - spiralPartHalfPoints) // target distance from center of side

					totalDistance += distanceFromCenter
					break
				}
			}

			return totalDistance
		}
	}
}

type Grid map[string]int

func stressTest(target int) {
	grid := make(Grid)

	grid["0x0"] = 1

	layer := 0
	x := 0
	y := 0

	for {
		layer += 1
		x += 1 // new layer

		edgeSize := layer * 2

		// right side up
		for i := 0; i < edgeSize; i++ {
			if i != 0 {
				y += 1
			}

			grid[gridKey(x, y)] = getSum(grid, x, y)
			check(grid, target, x, y)
		}

		// top side left
		for i := 0; i < edgeSize; i++ {
			x -= 1

			grid[gridKey(x, y)] = getSum(grid, x, y)
			check(grid, target, x, y)
		}

		// left side down
		for i := 0; i < edgeSize; i++ {
			y -= 1

			grid[gridKey(x, y)] = getSum(grid, x, y)
			check(grid, target, x, y)
		}

		// bottom side right
		for i := 0; i < edgeSize; i++ {
			x += 1

			grid[gridKey(x, y)] = getSum(grid, x, y)
			check(grid, target, x, y)
		}
	}
}

func check(grid Grid, target int, x int, y int) {
	if grid[gridKey(x, y)] > target {
		fmt.Println("solution", grid[gridKey(x, y)])

		os.Exit(0)
	}
}

/*
147  142  133  122   59
304    5    4    2   57
330   10    1    1   54
351   11   23   25   26
362  747  806--->   ...
*/

func getSum(grid Grid, x int, y int) int {
	ret := 0

	ret += get(grid, x+1, y+1)
	ret += get(grid, x+1, y)
	ret += get(grid, x+1, y-1)

	ret += get(grid, x-1, y+1)
	ret += get(grid, x-1, y)
	ret += get(grid, x-1, y-1)

	ret += get(grid, x, y+1)
	ret += get(grid, x, y-1)

	fmt.Printf("%2d %2d %d\n", x, y, ret)

	return ret
}

func get(grid Grid, x int, y int) int {
	val, ok := grid[gridKey(x, y)]

	if ok {
		return val
	}

	return 0
}

func gridKey(x int, y int) string {
	return fmt.Sprintf("%dx%d", x, y)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}

/*
17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---> ...
*/
