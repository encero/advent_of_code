package main

import (
  "fmt"
  "strings"
)

func main() {
  pass := strings.Split("hxbxxyzz", "")

  fmt.Printf("z -> %d\n", "z"[0]);

  fmt.Printf("%v -> hijklmmn\n", check(strings.Split("hijklmmn", "")))
  fmt.Printf("%v -> abbceffg\n", check(strings.Split("abbceffg", "")))
  fmt.Printf("%v -> abbcegjk\n", check(strings.Split("abbcegjk", "")))
  fmt.Printf("%v -> abcdefgh\n", check(strings.Split("abcdefgh", "")))
  fmt.Printf("%v -> abcdffaa\n", check(strings.Split("abcdffaa", "")))
  fmt.Printf("%v -> ghjaabcc\n", check(strings.Split("ghjaabcc", "")))


  for {
    inc(pass)
    //fmt.Printf("%s\n", strings.Join(pass, ""))
    if (check(pass)) {
      fmt.Printf("%s\n", strings.Join(pass, ""))
      break
    }
  }

}

func check(word []string) bool {
  for i := 0; i < len(word); i++ {
    char := word[i]
    if char == "i" || char == "o" || char == "l" {
      return false
    }
  }


  count := 0
  for i := 0; i < len(word) - 1; i++ {
    if (word[i] == word[i + 1]) {
      count ++
      i++
    }
  }

  ok := false
  for i := 0; i < len(word) - 2; i++ {
    if (word[i][0] == word[i+1][0] - 1 && word[i][0] == word[i+2][0] - 2) {
      ok = true
      break
    }
  }

  return count >= 2 && ok
}

func inc(word []string) {
  for index := len(word) - 1; index >= 0 ; index -- {
    char := word[index][0] + 1

    if (char > 122) {
      word[index] = "a"
    } else {
      word[index] = string(char)
      break
    }
  }
}
