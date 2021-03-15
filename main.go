package main

import (
	"fmt"
)

type Matrix [][]int

func NewMatrix(n int) Matrix {
	matrix := make(Matrix, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}
func (m Matrix) String() (result string) {
	result = ""
	for i := 0; i < len(m); i++ {
		result += fmt.Sprintln((m)[i])
	}
	return
}
func (m *Matrix) Size() int {
	return len(*m)
}
func (m *Matrix) Spiral() Matrix {
	up := 0
	down := m.Size() - 1
	left := 0
	right := m.Size() - 1
	i := 0
	j := 0
	maxSize := m.Size() * m.Size()

	for {
		if left++; left > right {
			break
		}
		for ; i < down; i++ {
			maxSize--
			(*m)[i][j] = maxSize
		}
		if down--; down < up {
			break
		}
		for ; j < right; j++ {
			maxSize--
			(*m)[i][j] = maxSize
		}
		if right--; right < left {
			break
		}
		for ; i > up; i-- {
			maxSize--
			(*m)[i][j] = maxSize
		}
		if up++; up > down {
			break
		}
		for ; j > left; j-- {
			maxSize--
			(*m)[i][j] = maxSize
		}
	}
	maxSize--
	(*m)[i][j] = maxSize
	return *m
}
func main() {
	matrix := NewMatrix(7)
	fmt.Println(matrix.Spiral())
}
