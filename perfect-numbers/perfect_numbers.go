package perfect

import (
	"errors"
)

var ErrOnlyPositive = errors.New("only positive number are accepted")

type Classification byte

const (
	ClassificationDeficient Classification = 'D'
	ClassificationPerfect                  = 'P'
	ClassificationAbundant                 = 'A'
)

func getAliquotSum(n int64) int64 {
	var s, i int64
	for i = 1; i <= n/2; i++ {
		if n%i == 0 {
			s += i
		}
	}
	return s
}

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return ClassificationAbundant, ErrOnlyPositive
	}
	aliquotSum := getAliquotSum(n)
	if aliquotSum == n {
		return ClassificationPerfect, nil
	} else if aliquotSum > n {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}
