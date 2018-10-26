package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		fmt.Println("file open failed", err)
		panic(err)
	}

	data, _  := ioutil.ReadAll(f)

	lines := strings.Split(string(data), "\n")

	if lines[len(lines) - 1] == "" {
		return lines[:len(lines) - 1]
	}

	return lines
}

func StringsToInts(lines []string) []int64 {
	out := make([]int64, len(lines))

	for i, v := range lines {
		tmp, _ := strconv.ParseInt(v, 10, 64)
		out[i] = tmp
	}

	return out
}
