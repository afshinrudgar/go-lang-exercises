package series

func All(n int, digits string) []string {
	is := len(digits) - n + 1
	res := make([]string, is)

	for i := 0; i < is; i++ {
		res[i] = digits[i : n+i]
	}

	return res
}

func UnsafeFirst(n int, digits string) string {
	return digits[0:n]
}
