package p12800_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p12k/p12800"
)

func TestSolve12865(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/12865")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4 7
6 13
4 8
3 6
5 12`,
			"14",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p12800.Solve12865(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
