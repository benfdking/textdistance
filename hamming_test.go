package textdistance

import (
	"fmt"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	t.Parallel()

	tts := []struct {
		ins  [2]string
		want float64
	}{
		{
			ins:  [2]string{"heath", "heath"},
			want: 0,
		},
		{
			ins:  [2]string{"river", "rover"},
			want: 1,
		},
		{
			ins:  [2]string{"1011101", "1001001"},
			want: 2,
		},
		{
			ins:  [2]string{"karolin", "kerstin"},
			want: 3,
		},
		{
			ins:  [2]string{"drummer", "dresser"},
			want: 3,
		},
		{
			ins:  [2]string{"flaw", "lawn"},
			want: 4,
		},
	}
	for _, tt := range tts {
		t.Run(fmt.Sprintf("%s", tt.ins), func(t *testing.T) {
			h := Hamming{}

			got := h.Distance(tt.ins[0], tt.ins[1])

			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
		})
	}
}

func BenchmarkHammingDistance(b *testing.B) {
	h := Hamming{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = h.Distance("karolin", "kerstin")
	}
}
