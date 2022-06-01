package p1000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
)

func TestSolve1085(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1085")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"6 2 10 3", "1"},
		{"1 1 5 5", "1"},
		{"653 375 1000 1000", "347"},
		{"161 181 762 375", "161"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p1000.Solve1085(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
