package etl

import "strings"

func Transform(in map[int][]string) (out map[string]int) {

	res := make(map[string]int)
	for score, ss := range in {
		for _, s := range ss {
			res[strings.ToLower(s)] = score
		}
	}

	return res
}
