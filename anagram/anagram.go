package anagram

import (
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func Detect(word string, candidates []string) []string {
	word = strings.ToLower(word)
	sortedWord := sortString(word)
	res := make([]string, 0)
	for _, o := range candidates {
		lo := strings.ToLower(o)
		if lo != word && sortString(lo) == sortedWord {
			res = append(res, o)
		}
	}
	return res
}
