package textdistance

import mapset "github.com/deckarep/golang-set"

// NewSymmetricalTversky returns a SymmetricalTversky structure with the StringToSet set to the default WordsToSet
func NewSymmetricalTversky(alpha, beta float64) SymmetricalTversky {
	return SymmetricalTversky{
		StringToSet: WordsToSet,
		Alpha:       alpha,
		Beta:        beta,
	}
}

type SymmetricalTversky struct {
	StringToSet func(s string) mapset.Set
	Alpha       float64
	Beta        float64
}

func (SymmetricalTversky) Maximum(s1, s2 string) float64 {
	return 1
}

func (t SymmetricalTversky) Similarity(s1, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}
	set1 := t.StringToSet(s1)
	set2 := t.StringToSet(s2)

	set1MinusSet2Cardinality := float64(set1.Difference(set2).Cardinality())
	set2MinusSet1Cardinality := float64(set2.Difference(set1).Cardinality())

	a := min(set1MinusSet2Cardinality, set2MinusSet1Cardinality)
	b := max(set1MinusSet2Cardinality, set2MinusSet1Cardinality)

	intersect := float64(set1.Intersect(set2).Cardinality())

	return intersect / (intersect + t.Beta*(t.Alpha*a+(1-t.Alpha)*b))
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}
