// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int64

const (
	// Pick values for the following identifiers used by the test program.
	NaT = 1 << iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// ShareWith should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	eqSides := 0
	if a == b {
		eqSides++
	}
	if a == c {
		eqSides++
	}
	if b == c {
		eqSides++
	}

	if eqSides == 3 {
		return Equ
	} else if eqSides == 1 {
		return Iso
	}
	return Sca
}
