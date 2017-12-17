package twofer

import "fmt"

func ShareWith(you string) string {
	if len(you) == 0 {
		you = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", you)
}
