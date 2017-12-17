package foodchain

import (
	"fmt"
	"strings"
)

type Animal struct {
	name, message, action string
}

var animals = []Animal{
	{name: "fly", message: "I don't know why she swallowed the fly. Perhaps she'll die."},
	{name: "spider", message: "It wriggled and jiggled and tickled inside her.", action: "wriggled and jiggled and tickled inside her"},
	{name: "bird", message: "How absurd to swallow a bird!"},
	{name: "cat", message: "Imagine that, to swallow a cat!"},
	{name: "dog", message: "What a hog, to swallow a dog!"},
	{name: "goat", message: "Just opened her throat and swallowed a goat!"},
	{name: "cow", message: "I don't know how she swallowed a cow!"},
	{name: "horse", message: "She's dead, of course!"},
}

func Verse(n int) string {

	lines := make([]string, 0)

	eaten := animals[n-1]
	lines = append(lines, fmt.Sprintf("I know an old lady who swallowed a %s.", eaten.name))
	if n > 1 {
		lines = append(lines, eaten.message)
	}
	if n == 8 {
		return strings.Join(lines, "\n")
	}

	for i := n - 2; i >= 0; i-- {
		eater := animals[i+1]
		eaten := animals[i]
		if eaten.action != "" {
			lines = append(lines, fmt.Sprintf("She swallowed the %s to catch the %s that %s.", eater.name, eaten.name, eaten.action))
		} else {
			lines = append(lines, fmt.Sprintf("She swallowed the %s to catch the %s.", eater.name, eaten.name))
		}
	}
	lines = append(lines, animals[0].message)
	return strings.Join(lines, "\n")
}

func Verses(a, b int) string {
	verses := make([]string, b-a+1)
	for i := a; i <= b; i++ {
		verses[i-a] = Verse(i)
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
