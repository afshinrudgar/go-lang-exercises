package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(ss []string) FreqMap {
	ch := make(chan FreqMap, len(ss))
	for _, s := range ss {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	res := FreqMap{}
	for i := 0; i < len(ss); i++ {
		fmap := <-ch
		for k, v := range fmap {
			res[k] += v
		}
	}

	close(ch)

	return res
}
