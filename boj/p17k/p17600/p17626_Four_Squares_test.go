package p17600_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p17k/p17600"
)

func TestSolve17626(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/17626")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"25", "1"},
		{"26", "2"},
		{"11339", "3"},
		{"34567", "4"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p17600.Solve17626(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
