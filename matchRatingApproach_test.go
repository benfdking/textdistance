package textdistance

import (
	"testing"
)

func TestMRAEncoding(t *testing.T) {
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
