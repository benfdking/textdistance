package textdistance

import (
	"fmt"
	"math"
	"testing"
)

func TestOverlap_Similarity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   [2]string
		want float64
	}{
		{
			in:   [2]string{"a b", "a b c d"},
			want: 1,
		},
		{
			in:   [2]string{"a b c d", "a b"},
			want: 1,
		},
		{
			in:   [2]string{"a b c d", "e f"},
			want: 0,
		},
		{
			in:   [2]string{"a b e g", "e f"},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s & %s", tt.in[0], tt.in[1]), func(t *testing.T) {
			o := NewOverlap()

			got := o.Similarity(tt.in[0], tt.in[1])

			if math.Abs(got-tt.want) > floatThreshold {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		})
	}
}

func BenchmarkOverlap_Similarity(b *testing.B) {
	const in1, in2 = "a b e g", "e f"
	o := Overlap{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		o.Similarity(in1, in2)
	}
}
