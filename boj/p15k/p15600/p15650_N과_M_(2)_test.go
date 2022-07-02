package p15600_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15600"
	"1d1go/utils"
)

func TestSolve15650(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15650")

	for _, tt := range []struct {
		give string
		want string
	}{
		{
			"3 1",
			`1
2
3
`,
		},
		{
			"4 2",
			`1 2
1 3
1 4
2 3
2 4
3 4
`,
		},
		{
			"4 4",
			`1 2 3 4
`,
		},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p15600.Solve15650(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
