package textdistance

import (
	"strings"

	mapset "github.com/deckarep/golang-set"
)

type Jaccard struct {
	StringToSet func(s string) mapset.Set
}

func NewJaccard() Jaccard {
	return Jaccard{StringToSet: WordsToSet}
}

func (Jaccard) Maximum(s1, s2 string) float64 {
	return 1
}

func (j Jaccard) Similarity(s1, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}
	set1 := j.StringToSet(s1)
	set2 := j.StringToSet(s2)
	return float64(set1.Intersect(set2).Cardinality()) / float64(set1.Union(set2).Cardinality())
}

func WordsToSet(s string) mapset.Set {
	set := mapset.NewSet()
	for _, token := range strings.Fields(s) {
		set.Add(token)
	}
	return set
}
