package p1800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"1d1go/boj/p1k/p1800"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve1874(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1874")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`8
4
3
6
8
7
5
2
1`,
			`+
+
+
+
-
-
+
+
-
+
+
-
-
-
-
-
`,
		},
		{
			`5
1
2
5
3
4`,
			"NO",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1800.Solve1874(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
