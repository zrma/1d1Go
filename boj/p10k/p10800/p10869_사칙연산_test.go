package p10800_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
)

func TestSolve10869(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10869")

	const (
		give = "7 3"
		want = `10
4
21
2
1
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	p10800.Solve10869(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
