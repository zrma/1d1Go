package p18200_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p18k/p18200"
)

func TestSolve18258(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/18258")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`15
push 1
push 2
front
back
size
empty
pop
pop
pop
size
empty
pop
push 3
empty
front`,
			`1
2
2
0
1
2
-1
0
1
-1
0
3
`,
		},
		{
			`2
front
back`,
			`-1
-1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p18200.Solve18258(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
