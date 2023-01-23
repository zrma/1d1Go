package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1774(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1774")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4 1
1 1
3 1
2 3
4 3
1 4`,
			"4.00",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1774(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
