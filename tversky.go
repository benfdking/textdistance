import mapset "github.com/deckarep/golang-set"

type Tversky struct {
	StringToSet func(s string) mapset.Set
	alpha       float64
	beta        float64
}

func NewTversky(alpha, beta float64) Tversky {
	return Tversky{
		StringToSet: WordsToSet,
		alpha:       alpha,
		beta:        beta,
	}
}

func (Tversky) Maximum(s1, s2 string) (float64, error) {
	return 1
}

func (t Tversky) Similarity(s1, s2 string) (float64, error) {
	if s1 == s2 {
		return 1, nil
	}
	set1 := j.StringToSet(s1)
	set2 := j.StringToSet(s2)
	union := set1.Union(set2)
	intersect := set1.Intersect(set2).Cardinality()
	bottton := t.alpha*union.Difference(set2).Cardinality() + t.beta*union.Difference(set1).Cardinality()
	return intersect / (intersect + bottom)
}



