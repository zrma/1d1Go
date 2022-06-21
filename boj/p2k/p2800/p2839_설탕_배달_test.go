package p2800_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2800"
	"1d1go/utils"
)

func TestSolve2839(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2839")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"0", "-1"},
		{"1", "-1"},
		{"2", "-1"},
		{"3", "1"},
		{"4", "-1"},
		{"5", "1"},
		{"6", "2"},
		{"7", "-1"},
		{"8", "2"},
		{"9", "3"},
		{"10", "2"},
		{"11", "3"},
		{"12", "4"},
		{"13", "3"},
		{"14", "4"},
		{"15", "3"},
		{"16", "4"},
		{"17", "5"},
		{"18", "4"},
		{"100", "20"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p2800.Solve2839(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
