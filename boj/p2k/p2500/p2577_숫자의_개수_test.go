package p2500_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2577(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2577")
	const (
		give = `150
266
427`
		want = `3
1
0
2
0
0
0
2
0
0
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p2500.Solve2577(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
