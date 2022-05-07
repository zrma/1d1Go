package codesprint5

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestExceedingTheSpeedLimit(t *testing.T) {
	for _, tt := range []struct {
		n    int32
		want string
	}{
		{100, "3000 Warning"},
		{140, "25000 License removed"},
		{85, "0 No punishment"},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			want := []string{tt.want}
			got, err := utils.GetPrinted(func() {
				exceedingTheSpeedLimit(tt.n)
			})
			assert.NoError(t, err)
			assert.Equal(t, want, got)
		})
	}
}
