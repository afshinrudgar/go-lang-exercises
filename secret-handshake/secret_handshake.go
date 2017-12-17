package secret

const (
	Wink        = 1 << iota
	DoubleBlink
	CloseEyes
	Jump
	Reverse
)

func Handshake(code uint) []string {
	reversed := code&Reverse != 0
	res := make([]string, 0)
	if code&Wink != 0 {
		res = append(res, "wink")
	}
	if code&DoubleBlink != 0 {
		res = append(res, "double blink")
	}
	if code&CloseEyes != 0 {
		res = append(res, "close your eyes")
	}
	if code&Jump != 0 {
		res = append(res, "jump")
	}
	if !reversed {
		return res
	}
	l := len(res)
	res2 := make([]string, l)
	for i, item := range res {
		res2[l-i-1] = item
	}
	return res2
}
