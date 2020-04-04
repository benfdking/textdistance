package textdistance

import (
	"math"
	"unicode/utf8"
)

// NewJaccard returns a Levenshtein structure
func NewLevenshtein () Levenshtein {
	return Levenshtein{}
}

type Levenshtein struct {
}

func (l Levenshtein) Distance(s1, b string) float64 {
	if len(s1) == 0 {
		return float64(utf8.RuneCountInString(b))
	}
	if len(b) == 0 {
		return float64(utf8.RuneCountInString(s1))
	}

	if s1 == b {
		return 0
	}

	rs1 := []rune(s1)
	rs2 := []rune(b)

	// swap to save some memory O(min(a,b)) instead of O(a)
	if len(rs1) > len(rs2) {
		rs1, rs2 = rs2, rs1
	}
	lenS1 := len(rs1)
	lenS2 := len(rs2)

	x := make([]int32, lenS1+1)
	// we start from 1 because index 0 is already 0.
	for i := int32(1); i < int32(lenS1); i++ {
		x[i] = i
	}

	// fill in the rest
	for i := 1; i <= int(lenS2); i++ {
		prev := int32(i)
		var current int32
		for j := 1; j <= int(lenS1); j++ {
			if rs2[i-1] == rs1[j-1] {
				current = x[j-1] // match
			} else {
				current = minUint(minUint(x[j-1]+1, prev+1), x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[lenS1] = prev
	}
	return float64(x[lenS1])
}

func minUint(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func (l Levenshtein) Minimum(s1, s2 string) float64 {
	l1, l2 := len(s1), len(s2)
	if l1 == l2 {
		return 0
	}
	return math.Abs(float64(l1) - float64(l2))
}

func (l Levenshtein) Maximum(s1, s2 string) float64 {
	l1, l2 := float64(len(s1)), float64(len(s2))
	return max(l1, l2)
}
