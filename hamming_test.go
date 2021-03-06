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
			h := NewHamming()

			got, err := h.Distance(tt.ins[0], tt.ins[1])

			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
			if err != nil {
				t.Errorf("expect empty error, got %+v", err)
			}
		})
	}
}

func BenchmarkHammingDistance(b *testing.B) {
	h := NewHamming()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := h.Distance("karolin", "kerstin")
		if err != nil {
			b.Fatal(err)
		}
	}
}
