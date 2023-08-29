package ky6

import "fmt"

func RunBubbleSort() {
	var (
		arr = []int{4, 9, 1, 3, 6, 2, 98, 4}
	)
	arr[3] = 5

	fmt.Println(arr)
	fmt.Println(DoubleSort(arr))
}

func DoubleSort(a []int) []int {
	var (
		sorted   = false
		isSwaped = false
	)

	for !sorted {
		isSwaped = false
		for ind := range a {
			if a[ind] > a[ind+1] {
				isSwaped = true
				a[ind], a[ind+1] = a[ind+1], a[ind]
			}
		}
		if !isSwaped {
			sorted = true
		}
	}
	return a
}
