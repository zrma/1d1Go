package p10800

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10814(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10814")

	//goland:noinspection SpellCheckingInspection
	const (
		give = `3
21 Junkyu
21 Dohyun
20 Sunyoung`
		want = `20 Sunyoung
21 Junkyu
21 Dohyun
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve10814(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
