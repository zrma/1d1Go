package p9000_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9000"
	"1d1go/utils"
)

func TestSolve9020(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9020")

	const (
		give = `5
1
4
8
10
16`
		want = `0 1
2 2
3 5
5 5
5 11
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p9000.Solve9020(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
