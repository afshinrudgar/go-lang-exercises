package lsproduct

import (
	"errors"
)

func getProduct(s string) (int, error) {
	r := 1
	var ci int
	for _, c := range s {
		ci = int(c - '0')
		if ci > 9 || c < 0 {
			return 0, errors.New("Invalid input")
		}
		r *= ci
	}
	return r, nil
}

func LargestSeriesProduct(in string, n int) (int, error) {
	if n == 0 {
		return 1, nil
	}
	if n < 1 || n > len(in) {
		return 0, errors.New("small input string")
	}
	lp := 0
	for i := 0; i <= len(in)-n; i++ {
		p, err := getProduct(in[i:i+n])
		if err != nil {
			return 0, err
		}
		if p > lp {
			lp = p
		}
	}
	return lp, nil
}
