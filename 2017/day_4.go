package main

import (
	"fmt"
	"github.com/encero/advent_of_code/tools"
	"strings"
)

func main() {
	lines := tools.ReadLines("./day_4.txt")

	fmt.Println("got", len(lines), "lines")

	valid := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		words := strings.Split(line, " ")

		isValid := true

		outer:
		for i, word := range words {
			for j := i + 1; j < len(words); j ++ {
				if word == words[j] || tools.AreAnagram(word, words[j]) {
					isValid = false
					break outer
				}
			}
		}

		if isValid {
			valid += 1
		}
	}

	fmt.Println("valid lines", valid)
}
