package pangram

import (
	"strings"
)

func IsPangram(in string) bool {

	visitedCharsCounter := 0
	visitedChars := make(map[rune]bool)
	for _, c := range strings.ToLower(in) {
		if c >= 'a' && c <= 'z' {
			_, ok := visitedChars[c]
			if ok {
				continue
			}
			visitedChars[c] = true
			visitedCharsCounter++
		}
	}

	return visitedCharsCounter == 26
}
