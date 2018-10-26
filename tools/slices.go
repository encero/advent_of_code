package tools

func StringIndexOf(arr []string, what string) int {
	for i, v := range arr {
		if v == what {
			return i
		}
	}

	panic("not in arr")
}
