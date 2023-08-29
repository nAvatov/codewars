package ky5

import (
	"fmt"
)

func RunFindAnnagrams() {
	fmt.Println(Anagrams("x", []string{"aabb", "abcd", "bbaa", "dada"}))
}

/*
Write a function that will find all the anagrams of a word from a list.
You will be given two inputs a word and an array with words.
You should return an array of all the anagrams or an empty array if there are none.

For example:
anagrams('abba', ['aabb', 'abcd', 'bbaa', 'dada']) => ['aabb', 'bbaa']

anagrams('racer', ['crazer', 'carer', 'racar', 'caers', 'racer']) => ['carer', 'racer']

anagrams('laser', ['lazing', 'lazy',  'lacer']) => []
*/

func Anagrams(word string, words []string) []string {
	var answer []string
	for _, v := range words {
		if len(v) != len(word) {
			continue
		}
		if SortLetters([]byte(v)) == SortLetters([]byte(word)) {
			answer = append(answer, v)
		}
	}
	return answer
}

func SortLetters(a []byte) string {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if rune(a[j]) < rune(a[i]) {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return string(a)
}
