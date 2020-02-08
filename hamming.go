package textdistance

type Hamming struct {
}

func (Hamming) Maximum(s1, s2 string) float64 {
	l1 := len(s1)
	l2 := len(s2)
	if l1 != l2 {
		panic("strings must be of same length")
	}
	return float64(l1)
}

func (Hamming) Minimum(s1, s2 string) float64 {
	return 0
}

// TODO Think about extending this with an error
func (Hamming) Distance(s1, s2 string) float64 {
	l1 := len(s1)
	l2 := len(s2)
	if l1 != l2 {
		panic("strings must be of same length")
	}
	r1 := []rune(s1)
	r2 := []rune(s2)
	var distance float64
	for i := range r1 {
		if r1[i] != r2[i] {
			distance++
		}
	}
	return distance
}

func (h Hamming) Similarity(s1, s2 string) float64 {
	return h.Distance(s1, s2) / h.Maximum(s1, s2)
}
