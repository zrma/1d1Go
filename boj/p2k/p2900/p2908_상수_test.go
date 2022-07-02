package p2900_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2900"
	"1d1go/utils"
)

func TestSolve2908(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2908")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"734 893", "437"},
		{"221 231", "132"},
		{"839 237", "938"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2900.Solve2908(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
