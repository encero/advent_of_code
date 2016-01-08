package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "io"
)

func badWords(word string) (bool) {
  bad  := []string{"ab", "cd", "pq", "xy"}

  for _, badWord := range bad {
    if strings.Contains(word, badWord) {
      return false
    }
  }

  return true
}

func repetition(word string) (bool) {
  var prev rune;

  for _, char := range word {
    if (prev == char) {
      return true
    }

    prev = char
  }

  return false
}

func vowels(word string) (bool) {
  count := 0;

  vowels := "aeiou";

  for _, char := range word {
    if strings.ContainsRune(vowels, char) {
      count += 1
    }

    if count >= 3 {
      break
    }
  }

  return count >= 3
}

func repetitionWithSkip(word string) (bool) {
  var prev, skip rune;

  for _, char := range word {
    if (prev == char) {
      return true
    }

    prev = skip
    skip = char
  }

  return false
}

func pairs(word string) (bool) {
  for i := 0 ; i < len(word) - 2; i++ {
    if strings.Contains(word[i + 2:], word[i: i + 2]) {
      return true
    }
  }

  return false
}

func main()  {
  in := bufio.NewReader(os.Stdin)

  nice := 0;
  nice2 := 0;

  for  {
    line, err := in.ReadString('\n')

//    if vowels(line) && repetition(line) && badWords(line) {
//      nice += 1
//    }

    if pairs(line) && repetitionWithSkip(line) {
      nice2 += 1
    }

    if (err == io.EOF) {
      break
    }
  }

  fmt.Printf("\n%d nice words\n", nice)
  fmt.Printf("%d nice2 words\n", nice2)
}
