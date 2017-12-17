package isogram

import (
  "unicode"
  "sort"
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

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

func IsIsogram(word string) bool {
  chars := make(map[rune]bool)
  // word = SortString(word)

  var exists bool
  for _, c := range word {
    c = unicode.ToUpper(c)

    _, exists = chars[c]
    if exists && unicode.IsLetter(c) {
      return false
    }
    chars[c] = true
  }

  return true
}
