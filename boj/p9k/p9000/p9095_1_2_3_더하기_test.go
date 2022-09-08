package p9000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9000"
	"1d1go/utils"
)

func TestSolve9095(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9095")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
4
7
10`,
			`7
44
274
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p9000.Solve9095(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
