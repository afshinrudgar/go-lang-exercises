package grains

import "errors"

func Square(s int) (uint64, error) {
	if s < 1 || s > 64 {
		return 0, errors.New("Invalid Parameter")
	}
	return 1 << uint(s-1), nil
}

func Total() uint64 {
	return 1<<uint(64) - 1
}
