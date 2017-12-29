package wordcount

import (
	"strings"
)

type Frequency map[string]int

func add(f Frequency, word string) {
	count, ok := f[word]
	if ok {
		f[word] = count + 1
	} else {
		f[word] = 1
	}
}

func WordCount(in string) Frequency {
	freq := Frequency{}
	currWord := make([]rune, 0)
	for _, c := range strings.ToLower(in) {
		if len(currWord) == 0 && c == '\'' {
			// start of a quote
			continue
		}
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '\'' {
			currWord = append(currWord, c)
			continue
		}
		l := len(currWord)
		if l > 0 {
			if currWord[l-1] == '\'' && l > 2 {
				currWord = currWord[:l-1]
			}
			word := string(currWord)
			add(freq, word)
		}
		currWord = make([]rune, 0)
	}
	if len(currWord) > 0 {
		add(freq, string(currWord))
	}
	return freq
}
