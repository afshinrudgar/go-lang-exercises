// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

func Abbreviate(s string) string {

	var res []byte

	var attach bool = true
	for _, p := range []byte(s) {
		if p == ' ' || p == '-' {
			attach = true
			continue
		}
		if attach {
			res = append(res, p&95)
			attach = false
		}
	}
	return string(res)
}
