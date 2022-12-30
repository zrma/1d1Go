package p17200_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p17k/p17200"
)

func TestSolve17298(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/17298")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4
3 5 2 7`,
			`5 7 7 -1 `,
		},
		{
			`4
9 5 4 8`,
			`-1 8 8 -1 `,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p17200.Solve17298(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
