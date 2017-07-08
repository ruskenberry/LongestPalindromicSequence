package lps

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"racecar", 7},
		{"hannah", 6},
	}
	for _, c := range tests {
		got := findLps(c.s)
		if got != c.want {
			t.Error("Longest Palindromic Subsequence for (%q) == %q, want %q", c.s, got, c.want)
		}
	}
}
