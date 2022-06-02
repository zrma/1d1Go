package p15500_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15500"
	"1d1go/utils"
)

func TestSolve15596(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15596")

	const (
		s = `10
1 10 4 9 2 3 8 5 7 6`
		want = "55"
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p15500.Solve15596(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}