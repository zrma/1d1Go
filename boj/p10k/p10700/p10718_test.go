package p10700_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10700"
	"1d1go/utils"
)

func TestSolve10718(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10718")

	//goland:noinspection SpellCheckingInspection
	want := []string{"강한친구 대한육군", "강한친구 대한육군"}
	got, err := utils.GetPrinted(func() {
		p10700.Solve10718()
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
