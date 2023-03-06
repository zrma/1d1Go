package p1100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1110(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1110")

	tests := []struct {
		give string
		want string
	}{
		{"26", "4"},
		{"55", "3"},
		{"1", "60"},
		{"0", "1"},
		{"71", "12"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1110(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
