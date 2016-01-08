package main

import (
  "fmt"
  "strconv"
  "strings"
)

func round(current []string) []string {
  count := 0
  var result []string
  var value, prev string = "" , ""

  for i := 0; i < len(current); i++ {
    value = current[i]

    if prev != value && prev != "" {
      result = append(result, strconv.Itoa(count))
      result = append(result, prev)
      count = 1
    } else {
      count++
    }

    prev = value
  }

  result = append(result, strconv.Itoa(count))
  result = append(result, prev)

  return result
}

func main() {
  word := strings.Split("3113322113", "")

  for i := 0; i < 50; i++ {
    word = round(word)
  }

  fmt.Printf("%s\n", len(word));
}
