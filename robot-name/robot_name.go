package robotname

import (
	"math/rand"
)

var poolSize int = 1000
var pool map[int][]string = make(map[int][]string, poolSize)

type Robot struct {
	name string
}

func hash(n string) int {
	var hv = []int{181, 191, 193, 197, 199}
	s := 0
	for idx, b := range []byte(n) {
		s += int(b) * hv[idx]
	}
	return s % poolSize
}

func generateRadnomName() string {
	cs := make([]byte, 5)
	cs[0] = byte(rand.Int31n(26)) + 'A'
	cs[1] = byte(rand.Int31n(26)) + 'A'
	cs[2] = byte(rand.Int31n(10)) + '0'
	cs[3] = byte(rand.Int31n(10)) + '0'
	cs[4] = byte(rand.Int31n(10)) + '0'
	return string(cs)
}

func getUniqueName() string {
	var name string
	var hashValue int
	var otherNames []string
	for {
		name = generateRadnomName()
		hashValue = hash(name)
		otherNames = pool[hashValue]
		for _, oName := range otherNames {
			if oName == name {
				continue
			}
		}
		pool[hashValue] = append(pool[hashValue], name)
		return name
	}
}

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = getUniqueName()
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}

// func New() *generator {
// 	return &generator{}
// }
