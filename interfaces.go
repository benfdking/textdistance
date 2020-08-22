package textdistance

// Distance generally tends to be a measure of the difference
type Distance interface {
	Distance(s1, s2 string) (float64, error)
}

// Minimum returns the quickly calculated theoretical minimum between two strings
type Minimum interface {
	Minimum(s1, s2 string) (float64, error)
}

// Maximum returns the theoretical maximum between two strings
type Maximum interface {
	Maximum(s1, s2 string) (float64, error)
}

// Similarity is a normalised measure of the distance between 0 and 1
type Similarity interface {
	Similarity(s1, s2 string) (float64, error)
}
