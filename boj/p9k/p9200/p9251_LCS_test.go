package p9200_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9200"
)

func TestSolve9251(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9251")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`ACAYKP
CAPCAK`,
			"4",
		},
		{
			`ABCDEF
GBCDEF`,
			"5",
		},
		{
			`ABCDEF
GBCDFE`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p9200.Solve9251(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
