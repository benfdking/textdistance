package textdistance

import "math"

type Levenshtein struct {
}

func (l Levenshtein) Distance(s1, s2 string) float64 {

}

func (l Levenshtein) Minimum(s1, s2 string) float64 {
	l1, l2 := len(s1), len(s2)
	if l1 == l2 {
		return 0
	}
	return math.Abs(float64(l1) - float64(l2))
}

func (l Levenshtein) Maximum(s1, s2 string) float64 {

}
