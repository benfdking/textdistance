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
			l := Levenshtein{}

			got := l.Minimum(tt.ins[0], tt.ins[1])

			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
		})
	}
}
