package summultiples

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}

	r := 0
	for b != 0 {
		r = a % b
		a = b
		b = r
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func SumMultiples(limit int, numbers ...int) int {

	if len(numbers) == 0 {
		return 0
	}

	res := 0
	for _, number := range numbers {
		for i := number; i < limit; i += number {
			res += i
		}
	}

	lcms := make([]int, 0)
	visited := make(map[int]bool)
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1 >= n2 {
				continue
			}
			l := lcm(n1, n2)
			_, ok := visited[l]
			if !ok {
				lcms = append(lcms, l)
			}
			visited[l] = true
		}
	}
	res -= SumMultiples(limit, lcms...)
	return res
}
