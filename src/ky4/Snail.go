package ky4

import "fmt"

func RunSnail() {
	anyQuadArray := [][]int{{}, {}, {}}

	// anyQuadArray := [][]int{{1, 2, 3, 4, 5},
	// 	{6, 7, 8, 9, 10},
	// 	{11, 12, 13, 14, 15},
	// 	{16, 17, 18, 19, 20},
	// 	{21, 22, 23, 24, 25}}

	// anyQuadArray := [][]int{{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9}}

	fmt.Println("Answer is: ", Snail(anyQuadArray))
}

// Если количество элементов - не четное, то последний элемент всегда имеет индекс ii, в противном случае ij.
// Можно рассматривать как итерационное прохождение краёв матрицы, начиная с элемента 00 и заканчивая элементом 10,
// если считать каждый следующий проход - новым массивом со своими индексами.

func Snail(snaipMap [][]int) []int {
	if len(snaipMap) == 0 || len(snaipMap[0]) == 0 {
		return []int{}
	}

	snail := GetSnail(snaipMap)

	return snail
}

func GetSnail(a [][]int) []int {
	answer := []int{}
	dir := true

	for startingIndex := 0; startingIndex <= int((len(a)-1)/2); startingIndex++ {
		answer = append(answer, TravelCorner(startingIndex, a, dir)...)
		answer = append(answer, TravelCorner(startingIndex, a, !dir)...)
	}

	return answer
}

func TravelCorner(start int, a [][]int, upper bool) []int {
	snailPart := []int{}
	i := start
	j := start

	if upper {

		for j = start; j < len(a)-start; j++ {
			snailPart = append(snailPart, a[i][j])
		}

		j--

		for i = start + 1; i < len(a)-start; i++ {
			snailPart = append(snailPart, a[i][j])
		}
		return snailPart
	}

	j = len(a) - start - 2
	i = j + 1

	for ; j >= start; j-- {
		snailPart = append(snailPart, a[i][j])
	}

	i--
	j++

	for ; i > start; i-- {
		snailPart = append(snailPart, a[i][j])
	}

	return snailPart
}
