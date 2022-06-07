package p1000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
)

func TestSolve1000(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1000")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"1 2", "3"},
		{"1 9", "10"},
		{"9 1", "10"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p1000.Solve1000(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}