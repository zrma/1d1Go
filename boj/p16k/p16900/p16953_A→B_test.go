package p16900_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p16k/p16900"
	"1d1go/utils"
)

func TestSolve16953(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/16953")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"2 162", "5"},
		{"4 42", "-1"},
		{"100 40021", "5"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p16900.Solve16953(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
