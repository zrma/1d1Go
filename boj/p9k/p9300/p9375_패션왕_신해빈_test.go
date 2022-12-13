package p9300_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9300"
)

func TestSolve9375(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9375")

	//goland:noinspection SpellCheckingInspection
	const (
		give = `2
3
hat headgear
sunglasses eyewear
turban headgear
3
mask face
sunglasses face
makeup face`
		want = `5
3
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p9300.Solve9375(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
