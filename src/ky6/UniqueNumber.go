package ky6

import "fmt"

func RunUniqueNumber() {
	array := []float32{0, 0, 0.55, 0, 0}
	fmt.Println(FindUnique(array))
}

func FindUnique(a []float32) float32 {
	var (
		unique  float32
		regular float32
	)

	if len(a) >= 3 {
		if a[0] == a[1] && a[1] == a[2] {
			regular = a[0]
			unique = LinearSearch(a, regular)
		} else if a[0] == a[1] && a[1] != a[2] {
			unique = a[2]
		} else if a[0] == a[2] && a[1] != a[2] {
			unique = a[1]
		} else if a[0] != a[2] && a[1] == a[2] {
			unique = a[0]
		}
	} else {
		return regular
	}

	return unique
}

func LinearSearch(arr []float32, regular float32) float32 {
	for i, v := range arr {
		if v != regular {
			return arr[i]
		}
	}
	return regular
}
