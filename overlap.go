package textdistance

import mapset "github.com/deckarep/golang-set"

// NewOverlap returns a Overlap structure with the StringToSet set to the default WordsToSet
func NewOverlap() Overlap {
	return Overlap{StringToSet: WordsToSet}
}

type Overlap struct {
	StringToSet func(string) mapset.Set
}

func (Overlap) Maximum(s1, s2 string) float64 {
	return 1
}

func (o Overlap) Similarity(s1, s2 string) float64 {
	if s1 == s2 {
		return 1
	}
	set1 := o.StringToSet(s1)
	set2 := o.StringToSet(s2)
	return float64(set1.Intersect(set2).Cardinality()) / min(float64(set1.Cardinality()), float64(set2.Cardinality()))
}
