package p1600_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1600"
)

func TestSolve1697(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1697")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"5 17", "4"},
		{"15 0", "15"},
		{"0 15", "6"},
		{"100000 0", "100000"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p1600.Solve1697(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
