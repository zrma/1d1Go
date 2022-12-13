package p3000_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p3k/p3000"
)

func TestSolve3053(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3053")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			"1",
			`3.141593
2.000000
`,
		},
		{
			"21",
			`1385.442360
882.000000
`,
		},
		{
			"42",
			`5541.769441
3528.000000
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p3000.Solve3053(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
