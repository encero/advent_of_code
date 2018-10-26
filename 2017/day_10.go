package main

import (
	"fmt"
)

func main() {
	input := "225,171,131,2,35,5,0,13,1,246,54,97,255,98,254,110"

	arr := make([]int, 256)
	for i := range arr {
		arr[i] = i
	}

	for _, v := range []int{17, 31, 73, 47, 23} {
		input += string(rune(v))
	}

	pointer := 0
	skip := 0
	for k := 0; k < 64; k++ {
		for _, v := range input {
			reverse(arr, pointer, int(v))

			pointer += int(v) + skip

			pointer = wrap(pointer, len(arr))
			skip += 1
		}
	}

	reduced := make([]int, 16)

	for i := range reduced {
		start := arr[i * 16]
		for j := 1; j < 16; j++ {
			start ^= arr[(i * 16) + j]
		}

		reduced[i] = start
	}

	for _, v := range reduced {
		fmt.Printf("%02x", uint16(v))
	}
	fmt.Println()



	fmt.Println(arr)
	fmt.Println(arr[0] * arr[1])

	_ = input
}

func reverse(arr []int, start, length int) {
	if length > (len(arr) - 1) {
		panic("len(arr) > length")
	}

	for i := (length / 2) - 1; i >= 0; i-- {
		a := wrap(start+length-i-1, len(arr))
		b := wrap(start+i, len(arr))

		tmp := arr[a]
		arr[a] = arr[b]
		arr[b] = tmp
	}
}

func wrap(x, max int) int {
	if x >= max {
		x = x - max

		return wrap(x, max)
	}

	return x
}
