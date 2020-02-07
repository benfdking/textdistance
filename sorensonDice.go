package textdistance

import mapset "github.com/deckarep/golang-set"

type SorensonDice struct {
	StringToSet func(s string) mapset.Set
}

func NewSorensonDice() SorensonDice {
	return SorensonDice{
		StringToSet: convertStringToSet,
	}
}

func (SorensonDice) Maximum(s1, s2 string) float {
	return 1
}

func (s SorensonDice) Similarity(s1, s2 string) float {
	if s1 == s2 {
		return 1.0
	}
	set1 := s.StringToSet(s1)
	set2 := s.StringToSet(s2)
	return 2 * float(set1.Intersect(set2).Cardinality()) / float((set1.Cardinality() + set2.Cardinality()))
}
