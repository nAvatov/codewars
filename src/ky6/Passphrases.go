package ky6

import (
	"bytes"
	"fmt"
)

func RunPassphrases() {
	fmt.Print(PlayPass("I LOVE YOU!!!", 1))
}

func PlayPass(s string, n int) string {
	var codedPass string
	codedPass = string(ShiftAndReplace(s, n))
	codedPass = string(EvenUp_OddDown(codedPass))
	codedPass = string(StringReverse(codedPass))
	return codedPass
}

func ShiftAndReplace(stringToChange string, shiftValue int) []byte { // Function shifts letter by a value of shiftValue and replace digits by they completment to 9
	return bytes.Map(func(r rune) rune {
		if r >= 97 && r <= 122 { // letters shift
			r += rune(shiftValue)
			r = ShiftByASCII(97, 122, r)
		} else if r >= 65 && r <= 90 {
			r += rune(shiftValue)
			r = ShiftByASCII(65, 90, r)
		} else if r >= 48 && r <= 57 { // digit replacement
			r = 47 + (58 - r)
		}
		return r
	}, []byte(stringToChange))
}

func ShiftByASCII(botBound rune, topBound rune, val rune) rune {
	if val > rune(topBound) {
		val = (botBound - 1) + (val - topBound)
	} else if val < botBound {
		val = (topBound + 1) - (botBound - val)
	}
	return val
}

func EvenUp_OddDown(stringToChange string) []byte {
	even := true
	return bytes.Map(func(r rune) rune {
		if even {
			if r >= 97 && r <= 122 {
				r -= 32
			}
			even = false
		} else {
			if r >= 65 && r <= 90 {
				r += 32
			}
			even = true
		}
		return r
	}, []byte(stringToChange))
}

func StringReverse(stringToReverse string) []byte {
	var reversedString []byte
	for i := len(stringToReverse) - 1; i >= 0; i-- {
		reversedString = append(reversedString, byte(stringToReverse[i]))
	}
	return reversedString
}
