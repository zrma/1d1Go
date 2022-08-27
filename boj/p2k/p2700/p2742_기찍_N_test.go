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

func TestSolve2742(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2742")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"1",
			`1
`,
		},
		{
			"2",
			`2
1
`,
		},
		{
			"3",
			`3
2
1
`,
		},
		{
			"4",
			`4
3
2
1
`,
		},
		{
			"5",
			`5
4
3
2
1
`,
		},
		{
			"6",
			`6
5
4
3
2
1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2700.Solve2742(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
