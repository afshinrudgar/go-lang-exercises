package pascal

func Triangle(n int) [][]int {

	if n <= 0 {
		return nil
	}

	res := make([][]int, n);
	res[0] = []int{1}

	for r := 2; r <= n; r++ {
		row := make([]int, r)
		preRow := res[r-2]

		row[0] = 1
		row[r-1] = 1

		for j := 1; j < r-1; j++ {
			row[j] = preRow[j-1] + preRow[j]
		}

		res[r-1] = row
	}

	return res
}
