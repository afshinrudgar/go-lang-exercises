package luhn

func Valid(word string) bool {

	toggle := false
	var v int
	var c byte
	var cint int
	var counter int = 0
	sum := 0

	for i := len(word) - 1; i >= 0; i-- {
		c = word[i]

		if c == ' ' {
			continue
		}

		cint = int(c)

		if cint < 48 || cint > 57 {
			return false
		}

		v = cint - 48
		counter++

		if toggle {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		sum += v
		toggle = !toggle

	}

	return counter > 1 && sum%10 == 0

}
