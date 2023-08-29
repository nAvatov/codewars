package ky6

import (
	"fmt"
	"strings"
)

func RunConsicutiveStrings() {
	//SortStringsByLen([]string{"zone", "abigail", "theta", "form", "libe", "zas"})
	fmt.Println(LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1))
}

func LongestConsec(a []string, k int) string {
	if k > len(a) {
		return ""
	}
	var (
		buf    string
		answer string
	)
	for i := 0; i <= len(a)-k; i++ {
		for j := i; j < (i + k); j++ {
			if !strings.Contains(buf, a[j]) {
				buf += a[j]
			}
		}
		if len(buf) > len(answer) {
			answer = buf
		}
		buf = ""
	}
	return answer
}
