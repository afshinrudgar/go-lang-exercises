package sublist

type Relation string

func isSublistOfFromIndex(a, b []int, idx int) bool {
	bLen := len(b)
  aLen := len(a)
  if aLen > bLen - idx {
    return false
  }
	for i := 0; i < aLen; i++ {
		if a[i] != b[idx] {
			return false
		}
    idx++
	}
	return true
}

func isSublistOf(a, b []int) bool {

	aLen, bLen := len(a), len(b)

	if aLen == 0 {
		return true
	}

  if aLen > bLen {
    return false
  }

	for idx := 0; idx < bLen-aLen+1; idx++ {
		if a[0] == b[idx] && isSublistOfFromIndex(a, b, idx) {
			return true
		}
	}

	return false
}

func Sublist(a, b []int) Relation {

	AonB := isSublistOf(a, b)
	BonA := isSublistOf(b, a)

	if AonB && BonA {
		return "equal"
	} else if AonB {
		return "sublist"
	} else if BonA {
		return "superlist"
	}
	return "unequal"
}
