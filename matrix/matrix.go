package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	Data       [][]int
	rNum, cNum int
}

func (m *Matrix) Set(r, c, v int) bool {
	if r < 0 || r >= m.rNum || c < 0 || c >= m.cNum {
		return false
	}
	m.Data[r][c] = v
	return true
}

func (m Matrix) Rows() [][]int {
	ret := make([][]int, m.rNum)
	for r := 0; r < m.rNum; r++ {
		ret[r] = make([]int, m.cNum)
		copy(ret[r], m.Data[r])
	}
	return ret
}

func (m Matrix) Cols() [][]int {
	ret := make([][]int, m.cNum)
	for c := 0; c < m.cNum; c++ {
		ret[c] = make([]int, m.rNum)
		for r := 0; r < m.rNum; r++ {
			ret[c][r] = m.Data[r][c]
		}
	}
	return ret
}

func New(m string) (*Matrix, error) {
	mat := &Matrix{}
	rows := strings.Split(m, "\n")
	mat.rNum = len(rows)
	for _, row := range rows {
		splitedRow := strings.Split(strings.TrimSpace(row), " ")
		if mat.cNum > 0 && mat.cNum != len(splitedRow) {
			return nil, errors.New("Invalid number of columns")
		}
		mat.cNum = len(splitedRow)
		irow := make([]int, len(splitedRow))
		for idx, rv := range splitedRow {
			num, err := strconv.Atoi(rv)
			if err != nil {
				return nil, err
			}
			irow[idx] = num
		}
		mat.Data = append(mat.Data, irow)
	}
	return mat, nil
}
