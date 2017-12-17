package diffsquares

func SumOfSquares(n int) int {
	s := 0
  for i := 1; i<=n; i++ {
    s += i*i
  }
	return s
}

func SquareOfSums(n int) int {
	t := (n + 1) * n / 2
	return t * t
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
