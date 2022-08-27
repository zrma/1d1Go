package p1100_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1100"
	"1d1go/utils"
)

func TestSolve1157(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1157")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{"Mississipi", "?"},
		{"zZa", "Z"},
		{"z", "Z"},
		{"zZ", "Z"},
		{"baaa", "A"},
		{"cddec", "?"},
		{`AA
BB`, "A"},
		{`AB
BB`, "?"},
		{"AB\nBB", "?"},
		{"AA\nBB", "A"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1100.Solve1157(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
