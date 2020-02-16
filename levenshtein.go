package textdistance

import "math"

type Levenshtein struct {
}

// TODO Move away from recursive to more efficient.
func (l Levenshtein) Distance(s1, s2 string) float64 {
	l1, l2 := len(s1), len(s2)

	// base case: empty strings
	if l1 == 0 {
		return float64(l2)
	}
	if l2 == 0 {
		return float64(l1)
	}

	var cost float64
	if s1[l1-1] == s2[l2-1] {
		cost = 0
	} else {
		cost = 1
	}
	return minimum(
		l.Distance(s1[:l1-1], s2)+1,
		l.Distance(s1, s2[:l2-1])+1,
		l.Distance(s1[:l1-1], s2[:l2-1])+cost,
	)
}

func minimum(a, b, c float64) float64 {
	minimum := a
	if minimum > b {
		minimum = b
	}
	if minimum > c {
		return c
	}
	return minimum
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
