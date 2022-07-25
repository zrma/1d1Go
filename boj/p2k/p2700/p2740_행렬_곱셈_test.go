package p2700_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
	"1d1go/utils"
)

func TestSolve2740(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2740")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 2
1 2
3 4
5 6
2 3
-1 -2 0
0 0 3`,
			`-1 -2 6
-3 -6 12
-5 -10 18
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2700.Solve2740(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
