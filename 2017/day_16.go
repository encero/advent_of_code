package main

import (
	"fmt"
	"github.com/encero/advent_of_code/tools"
	"strings"
)

func main() {
	input := strings.Split(tools.ReadLines("./day_16.txt")[0], ",")

	dancers := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p"}

	fmt.Println("count of dancers", len(dancers))

	cycles := make([]string, 0)

	count := 1000000000

	outer:
	for i := 0; i < count; i++ {
		name := strings.Join(dancers, "")

		for _, c := range cycles {
			if c == name {
				fmt.Println("result", cycles[count % i])
				break outer
			}
		}

		cycles = append(cycles, name)
		dancers = dance(input, dancers)

		fmt.Println(strings.Join(dancers, ""))
	}
}

func translate(arr []string, next []string, translation []int) {
	for i, v := range translation {
		next[v] = arr[i]
	}
}

func Switch(arr []string, a, b int ) {
	arr[a], arr[b] = arr[b], arr[a]
}

func dance(input, dancers []string) []string {
	for _, v := range input {
		switch v[0:1] {
		case "s":
			param :=len(dancers) - tools.To32(v[1:])

			dancers = append(append([]string{}, dancers[param:]...), dancers[:param]...)
		case "x":
			params := strings.Split(v[1:], "/")

			a := tools.To32(params[0])
			b := tools.To32(params[1])

			dancers[a], dancers[b] = dancers[b], dancers[a]
		case "p":
			params := strings.Split(v[1:], "/")

			a := tools.StringIndexOf(dancers, params[0])
			b := tools.StringIndexOf(dancers, params[1])

			dancers[a], dancers[b] = dancers[b], dancers[a]
		}
	}

	return dancers
}
