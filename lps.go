package lps

import "unicode/utf8"

func findLps(s string) int {
	b := []rune(s)
	n := utf8.RuneCountInString(s)
	c := make([][]int, n)

	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		for j := 0; j < n; j++ {
			c[i][j] = 1
		}
	}

	for cl := 2; cl <= n; cl++ {
		for i := 0; i < n-cl+1; i++ {
			var j = i + cl - 1

			if b[i] == b[j] && cl == 2 {
				c[i][j] = 2
			} else if b[i] == b[j] {
				c[i][j] = c[i+1][j-1] + 2
			} else {
				c[i][j] = max(c[i][j-1], c[i+1][j])
			}
		}
	}
	return c[0][n-1]
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
