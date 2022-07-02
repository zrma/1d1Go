package p7500_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p7k/p7500"
	"1d1go/utils"
)

func TestSolve7568(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/7568")

	const (
		give = `5
55 185
58 183
88 186
60 175
46 155`
		want = "2 2 1 2 5 "
	)

	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p7500.Solve7568(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
