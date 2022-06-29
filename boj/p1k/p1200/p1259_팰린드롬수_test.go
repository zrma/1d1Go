package p1200_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1200"
	"1d1go/utils"
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
	writer := utils.NewStringWriter()

	p1200.Solve1259(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
