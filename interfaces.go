package textdistance

type Distance interface {
	Distance(s1, s2 string) (float64, error)
}

type Minimum interface {
	Minimum(s1, s2 string) (float64, error)
}

type Maximum interface {
	Maximum(s1, s2 string) (float64, error)
}

type Similarity interface {
	Similarity(s1, s2 string) (float64, error)
}
