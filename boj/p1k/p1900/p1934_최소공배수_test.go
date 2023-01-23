package p1900

import (
	"bufio"
	_ "embed"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSolve1934(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1934")

	const (
		give = `4
1 45000
6 10
13 17
44999 45000`
		want = `45000
30
221
2024955000
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve1934(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}

//go:embed test_data/p1934_give.txt
var p1934give string

//go:embed test_data/p1934_want.txt
var p1934want string

func TestSolve1934_Performance(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(p1934give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	assert.Eventually(t, func() bool {
		Solve1934(reader, writer)
		return true
	}, time.Second, time.Millisecond*100, "시간 초과")

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, p1934want, got)
}
