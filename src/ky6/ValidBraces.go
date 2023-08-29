package ky6

import (
	"fmt"
)

func RunValidBraces() {
	var (
		answer = false
	)
	answer = ValidBraces("(}")
	fmt.Println(answer)
	//fmt.Print(IsClosedBracket(rune("}")))
}

func ValidBraces(s string) bool { // Approve that string isn't valid
	if len(s) == 0 {
		return true
	} else if len(s)%2 != 0 {
		return false
	} else {
		var (
			slicedString = ConvertToRuneArray(s)
			flag         = true
			valid        = true
			i            = 0
		)

		fmt.Println(slicedString)

		if IsClosedBracket(slicedString[0]) && IsOpenBracket(slicedString[len(slicedString)-1]) { // If first element is closed bracket and the last is opened
			return false
		} else {
			for flag {
				if len(slicedString) > 2 && i < len(slicedString)-2 && !IsClosedBracket(slicedString[i+1]) { // Moving forward while pair isnt found
					i++
				} else if IsOpenBracket(slicedString[i]) && IsClosedBracket(slicedString[i+1]) && Compare(slicedString[i], slicedString[i+1]) { // If we found a pair inside cut it form a slice
					slicedString = append(slicedString[:i], slicedString[i+2:]...) // Cutting slice
					i = 0
				} else {
					return false
				}
				if len(slicedString) == 0 {
					valid = true
					flag = false
					break
				} else if len(slicedString) == 1 {
					flag = false
					valid = false
					break
				}
			}
		}
		return valid
	}
}

func IsClosedBracket(bracket rune) bool {
	if bracket == 41 || bracket == 93 || bracket == 125 {
		return true
	} else {
		return false
	}
}

func IsOpenBracket(bracket rune) bool {
	if bracket == 40 || bracket == 91 || bracket == 123 {
		return true
	} else {
		return false
	}
}

func Compare(firstBracket rune, secondBracket rune) bool {
	if secondBracket-firstBracket <= 2 && secondBracket-firstBracket >= 1 {
		return true
	}
	return false
}

func ConvertToRuneArray(s string) []rune {
	arrayOfRunes := make([]rune, len(s))
	for i := range s {
		arrayOfRunes[i] = rune(s[i])
	}
	return arrayOfRunes
}
