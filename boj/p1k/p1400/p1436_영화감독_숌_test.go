package p1400_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1400"
	"1d1go/utils"
)

func TestSolve1436(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1436")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"1", "666"},
		{"2", "1666"},
		{"6", "5666"},
		{"7", "6660"},
		{"12", "6665"},
		{"13", "6666"},
		{"14", "6667"},
		{"16", "6669"},
		{"17", "7666"},
		{"19", "9666"},
		{"20", "10666"},
		{"25", "15666"},
		{"26", "16660"},
		{"27", "16661"},
		{"187", "66666"},
		{"500", "166699"},
		{"10000", "2666799"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1400.Solve1436(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
