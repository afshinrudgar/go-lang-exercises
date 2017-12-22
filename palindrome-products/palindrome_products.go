package palindrome

import (
	"strconv"
	"errors"
	"math"
)

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func reversed(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	return s == reversed(s)
}

func getFactorizations(n, fmin, fmax int) [][2]int {
	var d, frac float64
	fn := float64(n)
	ffmin := float64(fmin)
	ffmax := float64(fmax)
	facts := make([][2]int, 0)
	for i := ffmin; i < ffmax; i++ {
		d, frac = math.Modf(fn / i)
		if d < i {
			return facts
		}
		if frac == 0 && d >= ffmin && d <= ffmax {
			facts = append(facts, [2]int{int(i), int(d)})
		}
	}
	return facts
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax...")
	}

	var mn, mx int
	var mxProduct, mnProduct Product
	ffmin := fmin * fmin
	ffmax := fmax * fmax

	for mn = ffmin; mn < ffmax; mn++ {
		if !isPalindrome(mn) {
			continue
		}
		mnProduct = Product{Product: mn, Factorizations: getFactorizations(mn, fmin, fmax)}
		if len(mnProduct.Factorizations) > 0 {
			break
		}
	}

	if len(mnProduct.Factorizations) == 0 {
		return mnProduct, mxProduct, errors.New("no palindromes...")
	}

	for mx = ffmax; mx > mn; mx-- {
		if !isPalindrome(mx) {
			continue
		}
		mxProduct = Product{mx, getFactorizations(mx, fmin, fmax)}
		if len(mxProduct.Factorizations) > 0 {
			break
		}
	}

	if len(mxProduct.Factorizations) == 0 {
		return mnProduct, mxProduct, errors.New("no palindromes...")
	}

	return mnProduct, mxProduct, nil
}
