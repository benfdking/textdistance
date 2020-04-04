package textdistance

import mapset "github.com/deckarep/golang-set"

// NewSorensonDice returns a SorensonDice structure with the StringToSet set to the default WordsToSet
func NewSorensonDice() SorensonDice {
	return SorensonDice{
		StringToSet: WordsToSet,
	}
}

type SorensonDice struct {
	StringToSet func(s string) mapset.Set
}

func (SorensonDice) Maximum(s1, s2 string) float64 {
	return 1
}

func (s SorensonDice) Similarity(s1, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}
	set1 := s.StringToSet(s1)
	set2 := s.StringToSet(s2)
	return 2 * float64(set1.Intersect(set2).Cardinality()) / float64((set1.Cardinality() + set2.Cardinality()))
}
