package textdistance

type MRA struct {
}

func (m MRA) Minimum(s1, s2 string) float64 {
	sl := len(s1) + len(s2)
	switch {
	case sl <= 4:
		return 5
	case 4 < sl && sl <= 7:
		return 4
	case 7 < sl && sl <= 11:
		return 3
	case sl == 12:
		return 2
	default:
		panic("hello")
	}
}
