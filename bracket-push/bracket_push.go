package brackets

import (
	"errors"
)

type stack struct {
	items []rune
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) push(item rune) {
	s.items = append(s.items, item)
}

func (s *stack) pop() (item rune, err error) {
	l := s.length()
	if l == 0 {
		return ' ', errors.New("Empty Stack")
	}

	item = s.items[l-1]
	s.items = s.items[:l-1]
	err = nil
	return
}

func (s *stack) length() int {
	return len(s.items)
}

func contains(s []rune, i rune) (int, bool) {
	for idx, a := range s {
		if a == i {
			return idx, true
		}
	}
	return -1, false
}

func Bracket(in string) (bool, error) {
	start := []rune{'{', '[', '('}
	end := []rune{'}', ']', ')'}
	s := NewStack()

	for _, c := range in {
		sIdx, isStart := contains(start, c)
		if isStart {
			s.push(end[sIdx])
			continue
		}
		_, isEnd := contains(end, c)
		if isEnd {
			i, err := s.pop()
			if err != nil || c != i {
				return false, nil
			}
		}
	}

	if s.length() > 0 {
		return false, nil
	}

	return true, nil
}
