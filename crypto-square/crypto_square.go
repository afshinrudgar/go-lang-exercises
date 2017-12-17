package cryptosquare

import (
	"math"
	"strings"
)

func normalize(in string) (out string) {
	out = strings.ToLower(in)
	out = strings.Replace(out, " ", "", -1)
	out = strings.Replace(out, "#", "", -1)
	out = strings.Replace(out, "$", "", -1)
	out = strings.Replace(out, "!", "", -1)
	out = strings.Replace(out, "@", "", -1)
	out = strings.Replace(out, "%", "", -1)
	out = strings.Replace(out, "^", "", -1)
	out = strings.Replace(out, "&", "", -1)
	out = strings.Replace(out, ",", "", -1)
	out = strings.Replace(out, ".", "", -1)
	return
}

func Encode(in string) string {

	normalizedString := normalize(in)
	ln := len(normalizedString)

	if ln == 0 {
		return ""
	}

	s := math.Sqrt(float64(ln))
	cols := int(s*2+1) / 2
	rows := int(math.Ceil(s))
	res := ""
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			idx := col*rows + row
			if idx < ln {
				res += string(normalizedString[idx])
			}
		}
		res += " "
	}
	ln = len(res)
	return res[:ln-1]
}
