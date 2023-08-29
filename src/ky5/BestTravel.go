package ky5

import "fmt"

func RunBestTravel() {
	distances := []int{91, 74, 73, 85, 73, 81, 87}
	t := 331
	k := 4
	fmt.Println(ChooseBestSum(t, k, distances))

}

func ChooseBestSum(t, k int, ls []int) int {
	var (
		index_arr = []int{}
	)
	if t >= 0 && k >= 1 && k <= len(ls) && len(ls) >= 1 && ElementsCondition(ls) {
		return GenerateCombinations(ls, GenerateIndexArray(len(ls), index_arr), len(ls), k, t)
	} else {
		return -1
	}
}

func ElementsCondition(arr []int) bool {
	for _, v := range arr {
		if v < 0 {
			return false
		}
	}
	return true
}

func GenerateIndexArray(size int, a []int) []int {
	for i := 0; i < size; i++ {
		a = append(a, i+1)
	}
	return a
}

func GenerateCombinations(a []int, b []int, set int, subset int, border int) int {
	var (
		answer       = 0
		i            = subset - 1
		sum_distance = 0
	)
	answer = CalculateCombinationDistance(subset, border, answer, sum_distance, a, b)
	for {
		if b[i] < len(b)-subset+i+1 {
			b[i]++
			for j := i + 1; j <= subset-1; j++ { // make a combination
				b[j] = b[j-1] + 1
			}
			answer = CalculateCombinationDistance(subset, border, answer, sum_distance, a, b)
			i = subset
		}
		i--
		if i == -1 {
			if answer != 0 {
				return answer
			} else {
				return -1
			}
		}
	}
}

func CalculateCombinationDistance(subset int, border int, answer int, sum_distance int, a []int, b []int) int {
	for m := 0; m < subset; m++ { // indicate combination of indexes as an array values and sum it
		sum_distance += a[b[m]-1]
	}
	if sum_distance <= border && sum_distance > answer { // check and remember it
		answer = sum_distance
	}
	return answer
}
func Indicate(array []int) {
	for _, v := range array {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()
}
