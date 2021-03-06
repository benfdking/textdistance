package textdistance

import (
	"errors"
	"math"
	"strings"
)

// NewMRA returns a MRA structure
func NewMRA() MRA {
	return MRA{}
}

type MRA struct {
}

func (m MRA) Minimum(s1, s2 string) (float64, error) {
	e1, err := m.Encoding(s1)
	if err != nil {
		return 0, err
	}
	e2, err := m.Encoding(s2)
	if err != nil {
		return 0, err
	}

	sl := len(e1) + len(e2)
	switch {
	case sl <= 4:
		return 5, nil
	case 4 < sl && sl <= 7:
		return 4, nil
	case 7 < sl && sl <= 11:
		return 3, nil
	case sl == 12:
		return 2, nil
	default:
		return 0, errors.New("invalid length of strings, must be smaller than 12 combined")
	}
}

func (m MRA) Distance(s1, s2 string) (float64, error) {
	e1, err := m.Encoding(s1)
	if err != nil {
		return 0, err
	}
	e2, err := m.Encoding(s2)
	if err != nil {
		return 0, err
	}

	lengthDifference := math.Abs(float64(len(e1) - len(e2)))
	if lengthDifference > 3 {
		return 0, errors.New("encoded must have length difference less than 3")
	}

	//minimum := m.Minimum(e1, e2)

	//if math.Abs(len(encoded1) - len(encoded2)) > 3 {
	//
	//}
	//
	//min := m.Minimum(s1, s2)
	// TODO Finish

	return 0, errors.New("not fully implemented")
}

var (
	vowels = map[rune]bool{
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
)

// Encoding returns the encoded MRA string according to the match rating approach. Encoding follows the following steps:
//
// 1. Delete all vowels unless the vowel begins the word
// 2. Remove the second consonant of any double consonants present
// 3. Reduce codex to 6 letters by joining the first 3 and last 3 letters only
//
// From Wikipedia: https://en.wikipedia.org/wiki/Match_rating_approach
func (m MRA) Encoding(s string) (string, error) {
	s = strings.ToUpper(s)

	// step 1
	var removedVowels string
	for i := 0; i < len(s); i++ {
		if _, ok := vowels[rune(s[i])]; !ok || i == 0 {
			removedVowels += string(s[i])
		}
	}

	// step 2
	var removedDoubleConsonants string
	for i := 0; i < len(removedVowels)-1; i++ {
		if _, ok := vowels[rune(removedVowels[i])]; !ok {
			if _, ok := vowels[rune(removedVowels[i+1])]; !ok {
				removedDoubleConsonants += string(removedVowels[i])
			} else {
				i++
			}
		}
	}
	if _, ok := vowels[rune(removedVowels[len(removedVowels)-1])]; !ok {
		removedDoubleConsonants += string(removedVowels[len(removedVowels)-1])
	}

	// step 3
	if len(removedDoubleConsonants) > 6 {
		return removedDoubleConsonants[0:3] + removedDoubleConsonants[len(removedDoubleConsonants)-3:], nil
	}
	return removedDoubleConsonants, nil
}
