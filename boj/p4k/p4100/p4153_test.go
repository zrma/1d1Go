package p4100_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p4k/p4100"
	"1d1go/utils"
)

func TestSolve4153(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4153")

	const (
		s = `6 8 10
25 52 60
5 12 13
0 0 0`
		want = `right
wrong
right
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p4100.Solve4153(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}