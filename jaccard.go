package textdistance

import (
	mapset "github.com/deckarep/golang-set"
)

// NewJaccard returns a Jaccard structure with the StringToSet set to the default WordsToSet
func NewJaccard() Jaccard {
	return Jaccard{StringToSet: WordsToSet}
}

type Jaccard struct {
	StringToSet func(s string) mapset.Set
}

func (j Jaccard) Similarity(s1, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}
	set1 := j.StringToSet(s1)
	set2 := j.StringToSet(s2)
	return float64(set1.Intersect(set2).Cardinality()) / float64(set1.Union(set2).Cardinality())
}
