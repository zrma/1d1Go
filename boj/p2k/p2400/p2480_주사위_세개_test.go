package p2400_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2400"
	"1d1go/utils"
)

func TestSolve2480(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2480")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"3 3 6", "1300"},
		{"6 3 3", "1300"},
		{"2 2 2", "12000"},
		{"6 2 5", "600"},
		{"2 6 5", "600"},
		{"2 5 6", "600"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2400.Solve2480(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
