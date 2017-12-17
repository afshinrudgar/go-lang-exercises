package accumulate

func Accumulate(given []string, converter func(string) string) []string {
	res := make([]string, 0)
	for _, g := range given {
		res = append(res, converter(g))
	}
	return res
}
