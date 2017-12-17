package bob

// Hey should have a comment documenting it.
func Hey(remark string) string {

	var isEmpty bool = true
	var isAllCap bool = true
	var isAnyCap bool = false
	var isLastQM bool = false

	for _, c := range []byte(remark) {
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			continue
		}
		isEmpty = false

		if c >= 'a' && c <= 'z'{
			isAllCap = false
		}
		isAnyCap = isAnyCap || (c&32 == 0)
		isLastQM = c == '?'
	}

	if isEmpty {
		return "Fine. Be that way!"
	}

	if isAllCap && isAnyCap {
		return "Whoa, chill out!"
	}

	if isLastQM {
		return "Sure."
	}

	return "Whatever."
}
