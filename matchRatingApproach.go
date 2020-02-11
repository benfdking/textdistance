package textdistance

import (
	"strings"
)

type MRA struct {
}

func (m MRA) Minimum(s1, s2 string) float64 {
	sl := len(s1) + len(s2)
	switch {
	case sl <= 4:
		return 5
	case 4 < sl && sl <= 7:
		return 4
	case 7 < sl && sl <= 11:
		return 3
	case sl == 12:
		return 2
	default:
		panic("invalid length of strings, must be smaller than 12 combined")
	}
}

func (m MRA) Distance(s1, s2 string) float64 {
	//encoded1 := m.Encoding(s1)
	//encoded2 := m.Encoding(s2)

	//if math.Abs(len(encoded1) - len(encoded2)) > 3 {
	//
	//}
	//
	//min := m.Minimum(s1, s2)
	// TODO Finish

	return 0
}

var (
	vowels = map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
)

// Encoding returns the encoded MRA string according to the match rating approach. Encoding follows the following steps.
func (m MRA) Encoding(s string) string {
	lowerCaseString := strings.ToLower(s)

	var removedVowels string
	for _, r := range lowerCaseString {
		if _, ok := vowels[r]; !ok {
			removedVowels += string(r)
		}
	}

	var removedDoubleConsonants string
	for i := 0; i < len(removedVowels)-1; i++ {
		if _, ok := vowels[rune(removedVowels[i])]; !ok {
			if _, ok := vowels[rune(removedVowels[i+1])]; !ok {
				removedDoubleConsonants += string(removedVowels[i])
			}
		}
	}
	if _, ok := vowels[rune(removedVowels[len(removedVowels)-1])]; !ok {
		removedDoubleConsonants += string(removedVowels[len(removedVowels)-1])
	}

	l := len(removedDoubleConsonants)
	if l > 6 {
		return removedDoubleConsonants[0:3] + removedDoubleConsonants[l-3:l]
	}
	return strings.ToUpper(removedDoubleConsonants)
}
