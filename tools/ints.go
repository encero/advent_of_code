package tools

import "strconv"

func To32(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		panic(err)
	}

	return int(i)
}

func To64(s string) int64  {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return i
}
