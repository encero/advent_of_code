package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "io"
  "regexp"
  "strconv"
)

type Lights struct {
  status [1000][1000]int
}

type Point struct {
  x, y int
}

type Direction struct {
  from, to Point
  action string
}

func toInt(value string) (int) {
  converted, _ := strconv.Atoi(value)
  return converted
}

func parse(directions []string) (Direction) {
    dir := Direction{
    Point{toInt(directions[1]), toInt(directions[2])},
    Point{toInt(directions[3]), toInt(directions[4])},
    directions[0]}

    return dir
}

func count(lights Lights) int {
  count := 0
  for _, row := range lights.status {
    for _, cell := range row {
      count += cell
    }
  }

  return count
}

type do func(value int) int

func do_on(value int) int {
  return value + 1
}

func do_off(value int) int {
  value -= 1
  if value <= 0 {
    return 0
  } else {
    return value
  }
}

func do_toggle(value int) int {
  return value + 2
}

func doit(lights *Lights, dir Direction, f do) {
  for x:= dir.from.x; x <= dir.to.x; x++ {
    for y:= dir.from.y; y <= dir.to.y; y++ {
      lights.status[x][y] = f(lights.status[x][y])
    }
  }
}

func main()  {
  in := bufio.NewReader(os.Stdin)

  match, _ := regexp.Compile(`^(turn on|turn off|toggle) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)$`)

  lights := Lights{}

  for  {
    line, err := in.ReadString('\n')

    result := match.FindStringSubmatch(strings.Trim(line, "\n "))

    if len(result) == 6 {
      dir := parse(result[1:])

      switch dir.action {
      case "turn on":
        doit(&lights, dir, do_on)
      case "turn off":
        doit(&lights, dir, do_off)
      case "toggle":
        doit(&lights, dir, do_toggle)
      }

      fmt.Printf("%v\n", dir)
    }

    if (err == io.EOF) {
      break
    }
  }

  fmt.Printf("count: %d\n", count(lights))
  fmt.Printf("end\n");
}
