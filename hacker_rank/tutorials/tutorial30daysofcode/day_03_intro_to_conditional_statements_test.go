package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCond(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-conditional-statements/problem")

	for _, tt := range []struct {
		n    int32
		want string
	}{
		{3, weird},
		{24, notWeird},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got := cond(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
