package textdistance

type Distance interface {
	Distance(s1, s2 string) float64
}

type Minimum interface {
	Minimum(s1, s2 string) float64
}

type Maximum interface {
	Maximum(s1, s2 string) float64
}

type Similarity interface {
	Similarity(s1, s2 string) float64
}

func baseMaximum(strings []string) float64 {
	var maximum int
	for _, s := range strings {
		if len(s) > maximum {
			maximum = len(s)
		}
	}
	return float64(maximum)
}

func baseNormalizedDistance(distance, maximum float64) float64 {
	if maximum == 0 {
		return 1
	}
	return distance / maximum
}

func baseSimilarity(maximum, distance float64) float64 {
	return maximum - distance
}

func identical(strings []string) bool {
	// TODO Implement
	return false
}
