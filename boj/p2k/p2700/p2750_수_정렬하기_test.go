package p2700_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
	"1d1go/utils"
)

func TestSolve2750(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2750")

	const (
		give = `5
5
2
3
4
1`
		want = `1
2
3
4
5
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p2700.Solve2750(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
