package p15500

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve15552(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15552")

	arr2D := [][2]int{
		{1, 1}, {2, 3}, {3, 4}, {9, 8}, {5, 2},
	}
	want := []int{2, 5, 7, 17, 7}
	got := Solve15552(arr2D)
	assert.Equal(t, want, got)
}
