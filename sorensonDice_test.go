package textdistance

import "testing"

func BenchmarkSorensonDice_Similarity(b *testing.B) {
	const in1, in2 = "a b e g", "e f"
	o := NewSorensonDice()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := o.Similarity(in1, in2)
		if err != nil {
			b.Fatal(err)
		}
	}
}
