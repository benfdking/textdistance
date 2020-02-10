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

func (m MRA) Distance(s1, s2 string) float64{
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



const vowels = "aeiou"

func (m MRA) Encoding(s string) string {
	uppercaseVowels := strings.ToUpper(s)
	return uppercaseVowels
}
