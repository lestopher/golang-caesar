package main

import (
	"fmt"
	"math"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var frequencies = []float64{
	.0817, .0149, .0278, .0425, .1270, .0223, .0202, .0609, .0697, .0015, .0077, .0402, .0241,
	.0675, .0751, .0193, .0009, .0599, .0633, .0906, .0276, .0098, .0236, .0015, .0197, .0007,
}

func encrypt(s string, shift int) string {
	return rotate(s, shift)
}

func decrypt(s string, shift int) string {
	return rotate(s, -shift)
}

func rotate(s string, shift int) string {
	var sum float64
	shift %= 26
	byteArray := []byte(s)

	for pos, char := range byteArray {
		// XORed to make lower case, ignores spaces and numbers and other special characters
		char |= 0x20
		if 'a' <= char && char <= 'z' {
			byteArray[pos] = alphabet[(int(char)+shift)%26]
			sum += frequencies[(int(char)-int('a')+pos)%26]
		}
	}

	return string(byteArray)
}

func chi_square_sum(c string) float64 {
	length := len_chars_only(c)
	letterMap := make(map[rune]int)
	var sum float64

	// Get a count of how many times a letter occurs
	for _, char := range c {
		if 'a' <= char && char <= 'z' {
			_, ok := letterMap[char]
			if ok {
				letterMap[char] += 1
			} else {
				letterMap[char] = 1
			}
		}
	}

	for key, value := range letterMap {
		expectedFreq := frequencies[(int(key)-int('a'))%26] * float64(length)
		sum += math.Pow((float64(value)-expectedFreq), 2) / expectedFreq
	}

	return sum
}

func len_chars_only(c string) int {
	length := 0
	for _, char := range c {
		if 'a' <= char && char <= 'z' {
			length += 1
		}
	}
	return length
}

func main() {
	s := "ZL QRNE FVFGRE. V'Z URNQVAT GB GUR VAFGVGHGR. TBQ XABJF JUNG UNCCRARQ GB GUVF CYNPR ABJ. GUVF BYQ ONFGNEQ JVYY URYC HF. UR ZHFG URYC HF. --0321--"
	for i := 0; i < 26; i++ {
		decrytpedString := decrypt(s, i)
		fmt.Printf("%d\t%s\t%f\n", i, decrytpedString, chi_square_sum(decrytpedString))
	}
}
