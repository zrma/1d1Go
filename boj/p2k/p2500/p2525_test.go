package p2500_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
)

func TestSolve2525(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2525")

	for i, tt := range []struct {
		s    string
		want string
	}{
		{
			`14 30
20`,
			"14 50",
		},
		{
			`17 40
80`,
			"19 0",
		},
		{
			`23 48
25`,
			"0 13",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p2500.Solve2525(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
