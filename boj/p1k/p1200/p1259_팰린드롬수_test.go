package p1200

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1259(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1259")

	const (
		give = `121
1231
12421
5
0`
		want = `yes
no
yes
yes
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve1259(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
