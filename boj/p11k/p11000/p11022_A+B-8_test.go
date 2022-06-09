package p11000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11000"
	"1d1go/utils"
)

func TestSolve11022(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11022")

	const (
		s = `5
1 1
2 3
3 4
9 8
5 2`
		want = `Case #1: 1 + 1 = 2
Case #2: 2 + 3 = 5
Case #3: 3 + 4 = 7
Case #4: 9 + 8 = 17
Case #5: 5 + 2 = 7
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p11000.Solve11022(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}