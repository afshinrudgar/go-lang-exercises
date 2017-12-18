package sieve

func filter(in, out chan int, prime int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
}

func generate(in chan int, limit int) {
	for i := 2; i <= 2*limit; i++ {
		in <- i
	}
	close(in)
}

func Sieve(limit int) []int {

	res := make([]int, 0)

	g := make(chan int)
	go generate(g, limit)

	var prime int
	for {
		prime = <-g
		if prime > limit {
			break
		}
		res = append(res, prime)
		g2 := make(chan int)
		go filter(g, g2, prime)
		g = g2
	}

	return res

}
