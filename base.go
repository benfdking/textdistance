package textdistance

type Distance interface {
	Distance(strings string) float
}

func baseMaximum(strings []string) float {
	var maximum float
	for _, s := range strings {
		if len(s) > maximum {
			maximum = float(len(s))
		}
	}
	return maximum
}

func baseNormalizedDistance(distance, maximum float) float {
	if maximum == 0 {
		return 1
	}
	return distance / maximum
}

func baseSimilarity(maximum, distance float) float {
	return maximum - distance
}

func identical(strings []string) boolean {
	// TODO Implement
	return false
}
