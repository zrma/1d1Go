package p18900_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p18k/p18900"
)

func TestSolve18917(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/18917")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`12
1 3
1 1
1 4
3
4
1 1
3
4
2 1
2 4
3
4`,
			`8
6
9
7
4
2
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p18900.Solve18917(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
