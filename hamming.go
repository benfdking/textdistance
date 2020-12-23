package textdistance

import "errors"

// NewHamming returns a Hamming structure.
func NewHamming() Hamming {
	return Hamming{}
}

// Hamming structure incorporates methods for computing distance and similarity on Hamming.
type Hamming struct {
}

// Maximum returns the maximum value for Distance given two strings.
func (Hamming) Maximum(s1, s2 string) (float64, error) {
	l1 := len(s1)
	l2 := len(s2)
	if l1 != l2 {
		return 0, errors.New("strings must be of same length")
	}
	return float64(l1), nil
}

// Minimum returns the minimum value for Distance given two strings.
func (Hamming) Minimum(_, _ string) (float64, error) {
	return 0, nil
}

// TODO Think about extending this with an error
func (Hamming) Distance(s1, s2 string) (float64, error) {
	l1 := len(s1)
	l2 := len(s2)
	if l1 != l2 {
		return 0, errors.New("strings must be of same length")
	}
	r1 := []rune(s1)
	r2 := []rune(s2)
	var distance float64
	for i := range r1 {
		if r1[i] != r2[i] {
			distance++
		}
	}
	return distance, nil
}

func (h Hamming) Similarity(s1, s2 string) (float64, error) {
	d, err := h.Distance(s1, s2)
	if err != nil {
		return 0, err
	}
	m, err := h.Maximum(s1, s2)
	if err != nil {
		return 0, err
	}
	return d / m, nil
}
