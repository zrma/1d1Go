package p10900_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10900"
	"1d1go/utils"
)

func TestSolve10926(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10926")

	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		give string
		want string
	}{
		{"joonas", "joonas??!"},
		{"baekjoon", "baekjoon??!"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p10900.Solve10926(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
