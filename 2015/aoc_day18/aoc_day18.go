package main

import (
  "fmt"
  "bufio"
  "os"
  "io"
)

const SIDE_LENGTH = 100
var SIDES [][]int = [][]int{
  {-1, 0},
  {-1, -1},
  {0, -1},
  {1, 0},
  {1, 1},
  {0, 1},
  {-1, 1},
  {1, -1},
}

func makeLightsArray() [][]bool {
  lights := make([][]bool, SIDE_LENGTH)
  for i := 0; i < SIDE_LENGTH; i++ {
    lights[i] = make([]bool, SIDE_LENGTH)
  }

  return lights
}

func read() [][]bool {
  in := bufio.NewReader(os.Stdin)

  lights := makeLightsArray()

  lineNumber := 0
  for  {
    line, err := in.ReadString('\n')

    if (err == io.EOF || line == "") {
      break
    }

    for k, v := range line {
      if v == '#' {
        lights[lineNumber][k] = true
      } else if v == '.' {
        lights[lineNumber][k] = false
      }
    }

    lineNumber ++
  }

  return lights
}

func show(lights [][]bool) {
  fmt.Println("----")
  for _, arr := range lights {
    for _, v := range arr {
      if v {
        fmt.Printf("#")
      } else {
        fmt.Printf(".")
      }
    }
    fmt.Printf("\n")
  }
}

func count(lights [][]bool, x, y int) int {
  sum := 0

  for _, side := range SIDES {
    _x := x + side[0]
    _y := y + side[1]

    inBoundaries := _x >= 0 && _x < SIDE_LENGTH && _y >= 0 && _y < SIDE_LENGTH

    if inBoundaries {
      if (lights[_x][_y]) {
        sum ++
      }
    }
  }

  return sum
}

func step(lights, newLights [][]bool) {
  brokenLights(lights)

  for x := 0; x < SIDE_LENGTH; x++ {
    for y := 0; y < SIDE_LENGTH; y++ {
      newLights[x][y] = lights[x][y]
      cnt := count(lights, x, y)

      if lights[x][y] && (cnt != 2 && cnt != 3) {
        newLights[x][y] = false
      }

      if !lights[x][y] && cnt == 3 {
        newLights[x][y] = true
      }
    }
  }

  brokenLights(newLights)
}

func brokenLights(lights [][]bool) {
  lights[0][0] = true
  lights[0][SIDE_LENGTH - 1] = true
  lights[SIDE_LENGTH - 1][0] = true
  lights[SIDE_LENGTH - 1][SIDE_LENGTH - 1] = true
}

func main()  {
  lights := read()
  newLights := makeLightsArray()

  for i := 0; i <100; i++ {
    step(lights, newLights)
    lights, newLights = newLights, lights
  }

  sum := 0

  for _, arr := range lights {
    for _, v := range arr {
      if v {
        sum ++
      }
    }
  }

  fmt.Printf("%d\n", sum);
}
