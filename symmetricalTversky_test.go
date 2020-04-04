package textdistance

import "testing"

func BenchmarkSymmetricalTversky_Similarity(b *testing.B) {
	const in1, in2 = "a b e g", "e f"
	o := NewSymmetricalTversky(0.5, 0.5)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		o.Similarity(in1, in2)
	}
}
