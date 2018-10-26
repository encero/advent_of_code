package main

import (
	"fmt"
	"github.com/encero/advent_of_code/tools"
	"strings"
)

type Coord struct {
	X int
	Y int
	Z int
}

func (c *Coord) Add(x, y, z int) {
	c.X += x
	c.Y += y
	c.Z += z
}

func (c *Coord) Move(direction string) {
	switch direction {
	case "n":
		c.Add(0, 1, -1)
	case "s":
		c.Add(0, -1, 1)
	case "ne":
		c.Add(1, 0, -1)
	case "se":
		c.Add(1, -1, 0)
	case "nw":
		c.Add(-1, 1, 0)
	case "sw":
		c.Add(-1, 0, 1)
	default:
		panic("unkown direction " + direction)
	}
}

func (c1 *Coord) Distance(c2 Coord) int {
	return Max(Abs(c1.X-c2.X), Abs(c1.Y-c2.Y), Abs(c1.Z-c2.Z))
}

func main() {
	input := strings.Split(tools.ReadLines("./day_11.txt")[0], ",")

	//input = strings.Split("se,sw,se,sw,sw", ",")

	start := Coord{0, 0, 0}
	current := Coord{0, 0, 0}

	maxDistance := 0

	for _, v := range input {
		current.Move(v)
		if current.Distance(start) > maxDistance {
			maxDistance = current.Distance(start)
		}
	}

	fmt.Println("distance", current.Distance(start), "max distance", maxDistance)
}

func Max(a ... int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}

	return max
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}
