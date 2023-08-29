package ky6

import "fmt"

func RunPyramidArray() {
	fmt.Println(InitPyro(MakePyro(1)))
}

func MakePyro(n int) [][]uint {
	pyro := make([][]uint, n)

	for i := range pyro {
		pyro[i] = make([]uint, i+1)
	}

	return pyro
}

func InitPyro(pyro [][]uint) [][]uint {
	for i := range pyro {
		for j := range pyro[i] {
			pyro[i][j] = 1
		}
	}
	return pyro
}
