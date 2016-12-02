package main

import (
  "fmt"
  //"strconv"
)

// const TARGET = 25
// var LIST []int = []int{20, 15, 10, 5, 5}

const TARGET = 150
var LIST []int = []int{33,14,18,20,45,35,16,35,1,13,18,13,50,44,48,6,24,41,30,42}

func combinations(list []int, skip int) chan int {
  ch := make(chan int)

  fmt.Println(list[skip:])

  go func() {
    if skip < len(list) {
      for _, v := range list[skip:] {
        fmt.Printf("skip: %d v: %d\n", skip, v)
        ch <- v
      }
    }

    ch <- -1
  }()

  return ch
}

var MIN int = 100

func getMin() int {
  return MIN
}

func f(sum, skip int, count *int, used int) {
  if skip < len(LIST) {
    current := LIST[skip]

    if sum + current == TARGET {
      if used + 1 < MIN {
        MIN = used + 1
      }

      if (used + 1 == 4) {
        *count += 1
      }
    }

    if sum + current < TARGET {
      f(sum + current, skip + 1, count, used + 1)
    }

    f(sum, skip + 1, count, used)
  }
}
