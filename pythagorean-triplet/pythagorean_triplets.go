package pythagorean

import (
	"math"
)

type Triplet [3]int

func Range(min, max int) []Triplet {
	t := make([]Triplet, 0)
	var frac, c float64
	var ci int
	for a := min; a < max; a++ {
		for b := a + 1; b <= max; b++ {
			c = math.Sqrt(float64(a*a + b*b))
			c, frac = math.Modf(c)
			ci = int(c)
			if frac == 0 && ci >= min && ci <= max {
				t = append(t, Triplet{a, b, ci});
			}
		}
	}
	return t
}

func Sum(p int) []Triplet {
	res := make([]Triplet, 0)
	var c int
	for a := 1; a < p/2; a++ {
		for b := a + 1; b <= (p-a)/2; b++ {
			c = p - a - b
			if c*c == a*a+b*b {
				res = append(res, Triplet{a, b, c})
			}
		}
	}

	return res
}
