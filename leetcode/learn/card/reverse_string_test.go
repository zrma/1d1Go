package card

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	t.Log("https://leetcode.com/explore/learn/card/recursion-i/250/principle-of-recursion/1440/")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		given []byte
		want  []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		{[]byte{'A', 'B'}, []byte{'B', 'A'}},
	} {
		t.Run(fmt.Sprintf("%d/for", i), func(t *testing.T) {
			given := make([]byte, len(tt.given))
			copy(given, tt.given)

			reverseString(given)
			assert.Equal(t, tt.want, given)
		})

		t.Run(fmt.Sprintf("%d/recur", i), func(t *testing.T) {
			given := make([]byte, len(tt.given))
			copy(given, tt.given)

			reverseStringRecur(given)
			assert.Equal(t, tt.want, given)
		})
	}
}