package scrabble

import (
	"unicode"
  "sync"
)

var wordScore = map[rune]int{
	'A': 1, 'E': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'I': 1, 'D': 2, 'G': 2, 'B': 3, 'C': 3, 'F': 4, 'H': 4, 'K': 5, 'M': 3,
	'P': 3, 'V': 4, 'W': 4, 'Y': 4, 'J': 8, 'X': 8, 'Q': 10, 'Z': 10,
}

var wg sync.WaitGroup

func calcScore(c rune, out chan int) {
  defer wg.Done()
	out <- wordScore[unicode.ToUpper(c)]
}

func Score(word string) int {

  // not a great way to solve this problem
  // but it could be done using channels and synchronization
  
  l := len(word)
	tun := make(chan int, l)
	for _, c := range word {
    wg.Add(1)
		go calcScore(c, tun)
	}
  wg.Wait()
	score := 0
	for i := 0; i<l; i++ {
			score += <- tun
	}
	return score
}
