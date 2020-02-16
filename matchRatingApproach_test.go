package textdistance

import (
	"fmt"
	"math"
	"testing"
)

const floatThreshold = 0.0000000001

func TestMRA_Minimum(t *testing.T) {
	tests := []struct {
		in   [2]string
		want float64
	}{
		{
			in:   [2]string{"Byrne", "Boern"},
			want: 4,
		},
		{
			in:   [2]string{"Smith ", "Smyth"},
			want: 3,
		},
		// TODO Reimplement
		//	{
		//		in:   [2]string{"Catherine ", "Kathryn"},
		//		want: 3,
		//	},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s & %s", tt.in[0], tt.in[1]), func(t *testing.T) {
			mra := MRA{}

			got := mra.Minimum(tt.in[0], tt.in[1])

			if math.Abs(got-tt.want) > floatThreshold {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		})
	}
}

func TestMRA_Encoding(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "Byrne",
			want: "BYRN",
		},
		{
			in:   "Boern",
			want: "BRN",
		},
		{
			in:   "Smith",
			want: "SMTH",
		},
		{
			in:   "Smyth",
			want: "SMYTH",
		},
		{
			in:   "Catherine",
			want: "CTHRN",
		},
		{
			in:   "Kathryn",
			want: "KTHRYN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			mra := MRA{}

			got := mra.Encoding(tt.in)

			if got != tt.want {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}

func BenchmarkMRA_Encoding(b *testing.B) {
	const in = "boern"
	mra := MRA{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mra.Encoding(in)
	}
}
