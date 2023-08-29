package main

import (
	"fmt"
)

// func main() {
// 	s1 := "abcdef"
// 	s2 := "acf"
// 	fmt.Println(LCS(s1, s2))
// }

func LCS(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}

	result := []byte{}
	jStart := 0

	for i := 0; i < len(s2); i++ {
		for j := jStart; j < len(s1); j++ {
			if s2[i] == s1[j] {
				result = append(result, s2[i])
				jStart = j + 1
				break
			}
		}
	}
	fmt.Println(len(result))
	//test := string(result)
	fmt.Printf("%T\n", string(result))
	return string(result)
}
