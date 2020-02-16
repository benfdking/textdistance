package textdistance

import (
	"fmt"
	"testing"
)

func TestJaccard_Similarity(t *testing.T) {
	t.Parallel()

	tts := []struct {
		ins  [2]string
		want float64
	}{
		{
			ins:  [2]string{"c h", "a b c d e f g h"},
			want: 0.25,
		},
	}
	for _, tt := range tts {
		t.Run(fmt.Sprintf("%s, %s", tt.ins[0], tt.ins[1]), func(t *testing.T) {
			j := NewJaccard()

			got := j.Similarity(tt.ins[0], tt.ins[1])

			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
		})
	}
}

func BenchmarkJaccard_Similarity(b *testing.B) {
	const s1, s2 = "c h", "a b c d e f g h"

	j := NewJaccard()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		j.Similarity(s1, s2)
	}
}
