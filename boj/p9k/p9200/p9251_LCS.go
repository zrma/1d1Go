package p9200

import (
	"fmt"
)

func Solve9251(reader Reader, writer Writer) {
	var s0, s1 string
	_, _ = fmt.Fscan(reader, &s0, &s1)

	n0, n1 := len(s0), len(s1)
	dp := make([][]int, n0+1)
	for i := range dp {
		dp[i] = make([]int, n1+1)
	}

	for i := 1; i <= n0; i++ {
		for j := 1; j <= n1; j++ {
			if s0[i-1] == s1[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	_, _ = fmt.Fprint(writer, dp[n0][n1])
}
