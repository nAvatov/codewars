package ky6

import "fmt"

func RunMultiplicationTable() {
	fmt.Println(MultiplicationTable(3)[:])
}

func MultiplicationTable(n int) [][]int {
	return InitArray(CreateArray(n))
}

func CreateArray(size int) [][]int {
	array := make([][]int, size)
	for i := range array {
		array[i] = make([]int, size)
	}
	return array
}

func InitArray(array [][]int) [][]int {
	for i := range array {
		for j := range array {
			array[i][j] = (j + 1) * (i + 1)
		}
	}
	return array
}
