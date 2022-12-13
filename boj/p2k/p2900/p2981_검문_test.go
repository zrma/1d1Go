package p2900_test

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2900"
)

func TestSolve2981(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2981")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
6
34
38`,
			"2 4 ",
		},
		{
			`3
1
17
81`,
			"2 4 8 16 ",
		},
		{
			`5
5
17
23
14
83`,
			"3 ",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2900.Solve2981(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

//go:embed test_data/p2981_give.txt
var p2981give string

//go:embed test_data/p2981_want.txt
var p2981want string

func TestSolve2981_Performance(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(p2981give))
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	assert.Eventually(t, func() bool {
		p2900.Solve2981(reader, writer)
		return true
	}, time.Second, time.Millisecond*100, "시간 초과")

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	got = strings.TrimRight(got, " ") + "\n"
	assert.Equal(t, p2981want, got)
}
