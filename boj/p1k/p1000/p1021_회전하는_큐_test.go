package p1000_test

import (
	"bufio"
	"strings"
	"testing"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"

	"github.com/stretchr/testify/assert"
)

func TestSolve1021(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1021")

	for _, tt := range []struct {
		give string
		want string
	}{
		{
			`10 3
1 2 3`,
			"0",
		},
		{
			`10 3
2 9 5`,
			"8",
		},
		{
			`32 6
27 16 30 11 6 23`,
			"59",
		},
		{
			`10 10
1 6 3 2 7 9 8 4 10 5`,
			"14",
		},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1000.Solve1021(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
