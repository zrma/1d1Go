package p9300_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9300"
	"1d1go/utils"
)

func TestSolve9375(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9375")

	//goland:noinspection SpellCheckingInspection
	const (
		s = `2
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

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p9300.Solve9375(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
