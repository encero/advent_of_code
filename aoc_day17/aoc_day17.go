package main

import (
  "fmt"
)



func main()  {
  count := 0

  f(0, 0, &count, 0)

  fmt.Println("combinaci ", count)
  fmt.Println(getMin())
}
