package main

import (
  "fmt"
  "math"
)

func presents(house float64) float64 {
  sum := float64(0)

  root := math.Sqrt(house)

  for i  := float64(1); i <= root; i ++ {
    if house <= i * 50 && math.Mod(house, i) == 0 {
      sum += i * 11
    }

    if house/i != i && house <= house/i * 50 && math.Mod(house, i) == 0 {
      sum += house/i * 11
    }
  }

  return sum
}

func main()  {
  var house float64 = 1

  for {
    count := presents(float64(house))

    if count >= 36000000 {
      fmt.Println("---", house)
      break
    }

    if math.Mod(house, 1000) == 0 {
      fmt.Println(house, count)
    }
    house++
  }
}
