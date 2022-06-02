package p2400_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2400"
	"1d1go/utils"
)

func TestSolve2477(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2477")

	for i, tt := range []struct {
		s    string
		want string
	}{
		{
			`7
4 50
2 160
3 30
1 60
3 20
1 100`,
			"47600",
		},
		{
			`7
2 16
3 3
1 6
3 2
1 10
4 5`,
			"476",
		},
		{
			`1
1 2
3 2
2 1
4 1
2 1
4 1`,
			"3",
		},
		{
			`15
3 2
1 2
4 1
2 1
4 1
2 1`,
			"45",
		},
		{
			`1
4 10
2 10
4 20
2 20
3 30
1 30`,
			"700",
		},
		{
			`1
2 50
3 100
1 30
3 60
1 20
4 160`,
			"6200",
		},
		{
			`7
3 20
1 100
4 50
2 160
3 30
1 60`,
			"47600",
		},
		{
			`1
3 60
2 20
3 100
1 50
4 160
2 30`,
			"6800",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p2400.Solve2477(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
