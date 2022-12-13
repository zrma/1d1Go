package p2200_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2200"
)

func TestSolve2263(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2263")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
1 2 3
1 3 2`,
			"2 1 3 ",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2200.Solve2263(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
