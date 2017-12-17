package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	aLen := len(a)
	if aLen != len(b) {
		return -1, errors.New("disallow second strand longer")
	}

	d := 0
	for i := 0; i < aLen; i++ {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil
}
