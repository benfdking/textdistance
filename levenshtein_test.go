package textdistance

import (
	"fmt"
	"testing"
)

func TestLevenshtein_Minimum(t *testing.T) {
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
			ins:  [2]string{"dresser", "d"},
			want: 6,
		},
		{
			ins:  [2]string{"d", "dresser"},
			want: 6,
		},
	}
	for _, tt := range tts {
		t.Run(fmt.Sprintf("%s", tt.ins), func(t *testing.T) {
			l := NewLevenshtein()

			got, err := l.Minimum(tt.ins[0], tt.ins[1])

			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
			if err != nil {
				t.Errorf("expect empty error, got %+v", err)
			}
		})
	}
}

func TestLevenshtein_Distance(t *testing.T) {
	t.Parallel()

	tts := []struct {
		ins  [2]string
		want float64
	}{
		{
			ins:  [2]string{"flaw", "lawn"},
			want: 2,
		},
	}
	for _, tt := range tts {
		t.Run(fmt.Sprintf("%s", tt.ins), func(t *testing.T) {
			l := NewLevenshtein()

			got, err := l.Distance(tt.ins[0], tt.ins[1])

			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
			if err != nil {
				t.Errorf("expect empty error, got %+v", err)
			}
		})
	}
}

func BenchmarkLevenshtein_Distance(b *testing.B) {
	const s1, s2 = "flaw", "lawn"
	l := NewLevenshtein()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := l.Distance(s1, s2)
		if err != nil {
			b.Fatal(err)
		}
	}
}
